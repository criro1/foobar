package app

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

var ShutdownSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGKILL,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// RegisterShutdownBySignals - start listen system signal on app.Run() and call app.Shutdown() on got termination system signals
func (p *Pool) ShutdownOnSignal(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = ShutdownSignals
	}

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, signals...)

	go func() {
		defer close(ch)

		// listen system channel for closing app
		sig, ok := <-ch
		if ok {
			p.Logger.Info(`pool closing by signal`, zap.Stringer(`signal`, sig))
			p.Shutdown()
		}
	}()
}
