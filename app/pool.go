package app

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"go.uber.org/zap"
)

var (
	ErrNoClosersProvided          = errors.New(`no closers provided`)
	ErrReadinessForRunnerNotFound = errors.New(`readiness for runner not found`)
)

const (
	ExitCodeSuccess = 0
	ExitCodeError   = 1
)

func DoNothing() error {
	return nil
}

type (
	SignalChan     chan struct{}
	SignalReadChan <-chan struct{}

	Run   func() error
	Close func() error
	Ready func() <-chan struct{}

	Namer interface {
		Name() string
	}

	Readiness interface {
		Ready() <-chan struct{}
	}

	Runner interface {
		Run() error
	}

	Closer interface {
		Close() error
	}

	RunnerCloser interface {
		Runner
		Closer
	}

	Pool struct {
		runners     []Runner
		closers     []Closer
		readinesses []SignalReadChan

		readinessNames       map[string]int
		runnerAfterReadiness map[int]int

		Logger *zap.Logger

		ready       SignalChan // channel will be closed when all runners has been started
		finished    SignalChan // all runners finished
		shutdowning SignalChan // shut down in progress
		done        chan int   // exit code
		exitCode    int        // exit code

		upGroup    *sync.WaitGroup // wait for all runners started
		readyGroup *sync.WaitGroup // waiting for all readiness ready
		downGroup  *sync.WaitGroup // wait for all runners down

		onceShutdown    *sync.Once
		onceSetExitCode *sync.Once

		errs chan error
	}

	PoolOpt func(*Pool)
)

func (r Ready) Ready() <-chan struct{} {
	return r()
}

func (c Close) Close() error {
	return c()
}

func (r Run) Run() error {
	return r()
}

func WithLogger(l *zap.Logger) PoolOpt {
	return func(p *Pool) {
		p.Logger = l
	}
}

func NewPool(opts ...PoolOpt) *Pool {
	p := &Pool{
		runners:     make([]Runner, 0),
		closers:     make([]Closer, 0),
		readinesses: make([]SignalReadChan, 0),

		readinessNames:       make(map[string]int),
		runnerAfterReadiness: make(map[int]int),

		ready:       make(SignalChan),
		finished:    make(SignalChan),
		shutdowning: make(SignalChan),

		done:     make(chan int, 1), // buffer for exit code
		exitCode: -1,                // unknown exit code by default

		upGroup:    &sync.WaitGroup{},
		downGroup:  &sync.WaitGroup{},
		readyGroup: &sync.WaitGroup{},

		onceShutdown:    &sync.Once{},
		onceSetExitCode: &sync.Once{},

		errs: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(p)
	}

	if p.Logger == nil {
		p.Logger = zap.NewNop()
	}

	return p
}

func (p *Pool) Ready() <-chan struct{} {
	return p.ready
}

func (p *Pool) Done() <-chan int {
	return p.done
}

func (p *Pool) Runners() []Runner {
	return p.runners
}

func (p *Pool) Closers() []Closer {
	return p.closers
}

func (p *Pool) Readinesses() []SignalReadChan {
	return p.readinesses
}

func (p *Pool) AddRunner(r Runner) *Pool {
	p.runners = append(p.runners, r)
	return p
}

func (p *Pool) AddRunnerFn(r func() error) *Pool {
	return p.AddRunner(Run(r))
}

func (p *Pool) AddCloser(c Closer) *Pool {
	p.closers = append(p.closers, c)
	return p
}

func (p *Pool) AddCloserFn(c func() error) *Pool {
	return p.AddCloser(Close(c))
}

func (p *Pool) AddRunnerCloser(rc RunnerCloser) *Pool {
	return p.Add(rc)
}

func (p *Pool) Add(rc RunnerCloser) *Pool {
	p.AddRunner(rc)
	p.AddCloser(rc)

	if readiness, ok := rc.(Readiness); ok {
		p.AddReadiness(readiness)
	}
	return p
}

func (p *Pool) AddAfter(before Readiness, after RunnerCloser) error {
	newRunnerPos := len(p.runners)
	if err := p.addReadinessFor(before, newRunnerPos); err != nil {
		return err
	}

	p.Add(after)
	return nil
}

func Name(runner interface{}) string {
	if namer, ok := runner.(Namer); ok {
		return fmt.Sprintf(`%s:%v`, reflect.TypeOf(runner), namer.Name())
	} else {
		return fmt.Sprintf(`%s:%v`, reflect.TypeOf(runner), runner)
	}
}

func (p *Pool) addReadinessFor(before Readiness, runnerPos int) error {
	pos, ok := p.readinessNames[Name(before)]
	if !ok {
		return fmt.Errorf(`readiness = %s: %w`, Name(before), ErrReadinessForRunnerNotFound)
	}

	p.runnerAfterReadiness[runnerPos] = pos

	return nil
}

func (p *Pool) AddRunnerFnAfter(before Readiness, after Run) error {
	newRunnerPos := len(p.runners)
	if err := p.addReadinessFor(before, newRunnerPos); err != nil {
		return err
	}

	p.AddRunnerFn(after)
	return nil
}

func (p *Pool) AddReadiness(r Readiness) *Pool {
	p.readinesses = append(p.readinesses, r.Ready())
	p.readinessNames[Name(r)] = len(p.readinesses) - 1
	return p
}

func (p *Pool) runnerRun(pos int) {
	if before, ok := p.runnerAfterReadiness[pos]; ok {
		<-p.readinesses[before]
	}

	p.upGroup.Done() // done when runner just started

	defer func(i int) {
		if e := recover(); e != nil {
			err := fmt.Errorf(`%s panic: %w`, reflect.TypeOf(p.runners[pos]), fmt.Errorf(`%s`, e))
			p.errs <- err
		}

		p.Logger.Debug(`runner done`, zap.Int(`pos`, i))
		p.downGroup.Done()
	}(pos)

	err := p.runners[pos].Run()

	// runner returns error, exit code will be 1
	if err != nil {
		p.errs <- err
	}
}

func (p *Pool) readinessCheck() {
	readinessNum := len(p.readinesses)

	if readinessNum > 0 {
		p.Logger.Debug(`check readiness`, zap.Int(`num`, readinessNum))

		p.readyGroup.Add(readinessNum)

		for i := range p.readinesses {
			ready := p.readinesses[i]
			// each readiness separately observed
			go func() {
				<-ready
				p.readyGroup.Done()
			}()
		}

		p.readyGroup.Wait()
		p.Logger.Debug(`readiness done`, zap.Int(`num`, readinessNum))
	}
	// runner can start after some readiness, if it depends from it
	p.upGroup.Wait()

	close(p.ready)
}

func (p *Pool) shutdownOnErrors() {
	select {
	// stop caching if shut downing has been initiated
	case <-p.shutdowning:
		return

	// if error caught and pool not shut downing - shutdown initiated
	case err, ok := <-p.errs:
		if ok {
			p.Logger.Warn(`runner error`, zap.Error(err))
			p.setExitCode(ExitCodeError)
			p.Shutdown()
		}
		return

	case <-p.finished:
		p.Logger.Debug(`runners finished`)
		p.Shutdown()
		return
	}
}

func (p *Pool) Run() error {
	if len(p.closers) == 0 {
		return ErrNoClosersProvided
	}

	runnersNum := len(p.runners)
	p.Logger.Debug(`start runners`, zap.Int(`num`, runnersNum))

	p.upGroup.Add(runnersNum)
	p.downGroup.Add(runnersNum)
	p.errs = make(chan error, runnersNum)

	for pos := range p.runners {
		// each runner in separate goroutine
		go p.runnerRun(pos)
	}
	go p.readinessCheck()

	p.upGroup.Wait()

	// wait for finish
	go func() {
		p.downGroup.Wait()
		close(p.finished)
	}()

	// run catcher of error for shutdown application
	go p.shutdownOnErrors()

	return nil
}

// Shutdown pool
func (p *Pool) Shutdown() {
	// already shut downs¬
	if len(p.done) > 0 {
		return
	}

	// check shut downing. may be it will be
	select {
	case <-p.shutdowning:
		return
	default:
	}

	p.onceShutdown.Do(p.shutdown)
}

func (p *Pool) shutdown() {
	p.Logger.Debug(`pool shutdown in progress`)
	close(p.shutdowning) // closed  channel means that pool shutdown has been initiated

	// default exit code is success

	p.Logger.Debug(`call closers`, zap.Int(`num`, len(p.closers)))
	for i := len(p.closers) - 1; i >= 0; i -= 1 {
		if err := safelyCallCloser(p.closers[i]); err != nil {
			if p.exitCode == -1 {
				p.setExitCode(ExitCodeError)
			}

			p.Logger.Warn(`closer error`, zap.Error(err))
		}
	}

	p.Logger.Debug(`wait for runners finished`)
	p.downGroup.Wait()

	if p.exitCode == -1 {
		p.exitCode = ExitCodeSuccess
	}
	close(p.errs)
	p.Logger.Debug(`pool shut down`)

	p.done <- p.exitCode
}

func (p *Pool) setExitCode(code int) {
	p.Logger.Debug(`set exit code`, zap.Int(`code`, code))
	p.onceSetExitCode.Do(func() {
		p.exitCode = code
	})
}

func safelyCallCloser(c Closer) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(`%s`, e)
		}
	}()

	err = c.Close()
	return err
}
