package app

import (
	"time"
)

type (
	AppConfig struct {
		Listen  Listen
		Debug   *Debug
		Metrics *Metrics
		Logging *Logging
		Tracing *Tracing
	}

	// Listen configuration for address and port
	Listen struct {
		HTTP string
		GRPC string
	}

	ListenDefaults struct {
		HTTP    string
		HTTPEnv string
		GRPC    string
		GRPCEnv string
	}

	Debug struct {
		Listen string `validate:"required"`
		Path   string
	}

	Metrics struct {
		Path             string
		HTTP             bool
		GRPC             bool
		GRPCConnectivity time.Duration
	}

	Logging struct {
		HTTP *LoggingHTTP
	}

	LoggingHTTP struct {
		Request      bool
		RequestBody  bool
		ResponseBody bool
	}

	Tracing struct {
		Type     string
		Endpoint string
	}
)
