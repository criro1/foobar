package app

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// EnvVarLogDevelopment usage: LOG_DEV={any string except `0` and `false`} - enables log development mode
	EnvVarLogDevelopment = `LOG_DEV`

	// EnvVarLogLevel usage: LOG_LEVEL=info - enables info log level (also `warning`, `debug`, `error` accepted)
	EnvVarLogLevel = `LOG_LEVEL`
)

var (
	Config zap.Config

	// from zap/global.go
	_globalMu sync.RWMutex
	_globalL  *zap.Logger
	_globalS  *zap.SugaredLogger

	envLogDev   string
	envLogLevel string
)

func init() {
	envLogDev, _ = os.LookupEnv(EnvVarLogDevelopment)
	envLogLevel, _ = os.LookupEnv(EnvVarLogLevel)

	Config = configFromEnv(envLogDev, envLogLevel)

	_globalL, _ = Config.Build()
	_globalS = _globalL.Sugar()

}

func configFromEnv(logDev string, logLevel string) zap.Config {
	var config zap.Config

	if len(logDev) > 0 && logDev != `0` && logDev != `false` {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
		config.DisableStacktrace = true
	}

	if logLevel != `` {
		var l = new(zapcore.Level)
		err := l.UnmarshalText([]byte(envLogLevel))
		if err == nil {
			config.Level = zap.NewAtomicLevelAt(*l)
		}
	}

	return config
}

func Named(name string) *zap.SugaredLogger {
	ReplaceGlobals(L().Named(name))
	return S()
}

func DefaultDev() {
	// no explicit env var LOG_DEV and current config is not dev
	if envLogDev == `` && !Config.Development {
		_globalMu.Lock()
		Config = configFromEnv(`true`, envLogLevel)
		logger, _ := Config.Build()
		_globalMu.Unlock()

		ReplaceGlobals(logger)
	}
}

func ReplaceGlobals(logger *zap.Logger) func() {
	_globalMu.Lock()
	prev := _globalL
	_globalL = logger
	_globalS = logger.Sugar()
	_globalMu.Unlock()
	return func() { ReplaceGlobals(prev) }
}

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
// It's safe for concurrent use.
func L() *zap.Logger {
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}

// S returns the global SugaredLogger, which can be reconfigured with
// ReplaceGlobals. It's safe for concurrent use.
func S() *zap.SugaredLogger {
	_globalMu.RLock()
	s := _globalS
	_globalMu.RUnlock()
	return s
}
