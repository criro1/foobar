package app

import (
	"fmt"
	"net"
	"net/http"
	"reflect"

	"go.uber.org/zap"
)

type (
	HttpServer struct {
		Server  *http.Server
		Enabled bool
		Mux     *http.ServeMux

		name         string
		attrs        map[string]string
		rootRedirect string
		swaggers     [][]byte

		listener net.Listener
		logger   *zap.Logger
	}

	HTTPServerOpt func(*HttpServer) error
)

func WithName(name string) HTTPServerOpt {
	return func(s *HttpServer) error {
		s.name = name
		return nil
	}
}

func WithHTTPRequestLogging(logRequestBody, logResponseBody bool) HTTPServerOpt {
	return func(h *HttpServer) error {
		h.attrs[`request logging`] = fmt.Sprintf(
			`on, request body = %t, response body=%t`, logRequestBody, logResponseBody)

		h.Server.Handler = HTTPHandlerMiddleware(
			h.Mux, h.logger, logRequestBody, logResponseBody)
		return nil
	}
}

// NewDisabledHTTP used when only grpc endpoint config is provided
func NewDisabledHTTP(logger *zap.Logger) *HttpServer {
	return &HttpServer{
		Enabled: false,
		logger:  logger,
	}
}

func NewHTTP(addr string, logger *zap.Logger, opts ...HTTPServerOpt) (*HttpServer, error) {
	listener, err := net.Listen(`tcp`, addr)
	if err != nil {
		return nil, fmt.Errorf(`app http listener: %w`, err)
	}

	mux := http.NewServeMux()

	srv := &HttpServer{
		Server:  &http.Server{Handler: mux},
		Mux:     mux,
		Enabled: true,

		attrs:    make(map[string]string),
		listener: listener,
		logger:   logger,
	}

	for _, opt := range opts {
		if err = opt(srv); err != nil {
			return nil, fmt.Errorf(`app opt: %w`, err)
		}
	}

	if srv.logger == nil {
		srv.logger = zap.NewNop()
	}

	return srv, nil
}

func (h *HttpServer) SetRedirectPath(redirectPath string) {
	h.rootRedirect = redirectPath
}

func (h *HttpServer) GetRedirectPath() string {
	return h.rootRedirect
}

func (h *HttpServer) Run() error {
	if !h.Enabled {
		h.logger.Info(`http endpoint is not configured`)
		return nil
	}
	var attrs []zap.Field

	if h.name != `` {
		attrs = append(attrs, zap.String(`name`, h.name))
	}

	attrs = append(attrs,
		zap.Stringer(`addr`, h.listener.Addr()),
		zap.String(`listener`, reflect.TypeOf(h.listener).String()),
		zap.String(`handler`, reflect.TypeOf(h.Server.Handler).String()))

	for k, v := range h.attrs {
		attrs = append(attrs, zap.String(k, v))
	}

	h.logger.Info(`http run`, attrs...)

	err := h.Server.Serve(h.listener)
	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (h *HttpServer) Close() error {
	if !h.Enabled {
		return nil
	}
	h.logger.Info(`http closing`, zap.Stringer(`addr`, h.listener.Addr()))
	err := h.Server.Close()
	h.logger.Info(`http closed`, zap.Stringer(`addr`, h.listener.Addr()))

	return err
}
