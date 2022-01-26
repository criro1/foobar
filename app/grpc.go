package app

import (
	"context"
	"io"
	"net"
	"reflect"
	"strings"
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcmiddlewarerecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcmiddlewaretags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/s7techlab/cckit/gateway"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	channelzsvc "google.golang.org/grpc/channelz/service"
)

const (
	ChannelZMethod = `grpc.channelz.v1.Channelz`

	// GRPCDefaultGracefulStopTimeout - period to wait result of grpc.GracefulStop
	// after call grpc.Stop
	GRPCDefaultGracefulStopTimeout = 5 * time.Second
)

type (
	GrpcServer struct {
		Server *grpc.Server

		addr     string
		listener net.Listener
		name     string
		attrs    map[string]string

		logger          *zap.Logger
		ready           chan struct{}
		localConnection *grpc.ClientConn

		GracefulStopTimeout time.Duration
	}

	GRPCServerOpt func(*grpcServerOptions) error

	grpcServerOptions struct {
		streamInterceptors []grpc.StreamServerInterceptor
		unaryInterceptors  []grpc.UnaryServerInterceptor
	}
)

// UnaryServerInterceptor adds unary interceptor to server
func UnaryServerInterceptor(interceptor grpc.UnaryServerInterceptor) GRPCServerOpt {
	return func(options *grpcServerOptions) error {
		options.unaryInterceptors = append(options.unaryInterceptors, interceptor)

		return nil
	}
}

func NewGrpc(addr string, logger *zap.Logger, opts ...GRPCServerOpt) (*GrpcServer, error) {
	listener, err := net.Listen(`tcp`, addr)
	if err != nil {
		return nil, err
	}

	// add ChannelZ method to ignore
	logOptions := []grpczap.Option{
		grpczap.WithDecider(func(fullMethod string, err error) bool {
			if strings.Contains(fullMethod, ChannelZMethod) && err == nil {
				return false
			}
			return true
		}),
	}

	// Default configuration options
	defaultOpts := grpcServerOptions{
		streamInterceptors: []grpc.StreamServerInterceptor{
			grpcmiddlewaretags.StreamServerInterceptor(),
			grpczap.StreamServerInterceptor(logger, logOptions...),
			grpcmiddlewarerecovery.StreamServerInterceptor(),
		},
		unaryInterceptors: []grpc.UnaryServerInterceptor{
			grpcmiddlewaretags.UnaryServerInterceptor(),
			grpczap.UnaryServerInterceptor(logger, logOptions...),
			grpcmiddlewarerecovery.UnaryServerInterceptor(),
		},
	}

	// Apply custom options
	for _, opt := range opts {
		if err = opt(&defaultOpts); err != nil {
			return nil, err
		}
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(defaultOpts.streamInterceptors...)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(defaultOpts.unaryInterceptors...)),

		// todo: enable only if metrics Enabled
		grpc.StatsHandler(&ocgrpc.ServerHandler{
			IsPublicEndpoint: false,
			StartOptions: trace.StartOptions{
				SpanKind: trace.SpanKindServer,
			},
		}),
	)

	srv := &GrpcServer{
		Server: grpcServer,

		addr:     addr,
		listener: listener,
		attrs:    make(map[string]string),
		logger:   logger,
		ready:    make(chan struct{}),

		GracefulStopTimeout: GRPCDefaultGracefulStopTimeout,
	}

	return srv, nil
}

func (g *GrpcServer) Name() string {
	return g.name
}

func (g *GrpcServer) RegisterServices(defs []gateway.ServiceDef) {
	for _, def := range defs {
		g.logger.Debug(`grpc service register`,
			zap.String(`name`, def.Desc.ServiceName),
			zap.Stringer(`service`, reflect.TypeOf(def.Service)))

		g.Server.RegisterService(def.Desc, def.Service)
	}
}

func (g *GrpcServer) RegisterServicesToMux(defs []gateway.ServiceDef) (*runtime.ServeMux, error) {
	// register gRPC services
	g.RegisterServices(defs)

	serveMux := runtime.NewServeMux(runtime.WithMarshalerOption(
		//runtime.MIMEWildcard, &runtime.JSONBuiltin{}),
		//runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
			runtime.MIMEWildcard, &JSONLocal{JSONPb: &runtime.JSONPb{OrigName: true, EmitDefaults: true}, JSONBuiltin: &runtime.JSONBuiltin{}}),
		runtime.WithForwardResponseOption(HTTPResponseModifier),
		runtime.WithProtoErrorHandler(HTTPResponseErrorModifier))

	opts := []grpc.DialOption{grpc.WithInsecure()}
	for _, def := range defs {
		if err := def.HandlerFromEndpointRegister(context.Background(), serveMux, g.addr, opts); err != nil {
			return nil, err
		}
	}

	return serveMux, nil
}

type JSONLocal struct {
	JSONPb *runtime.JSONPb
	JSONBuiltin *runtime.JSONBuiltin
}

func (j *JSONLocal) Marshal(v interface{}) ([]byte, error) {
	return j.JSONPb.Marshal(v)
}

func (j *JSONLocal) Unmarshal(data []byte, v interface{}) error {
	return j.JSONPb.Unmarshal(data, v)
}

func (j *JSONLocal) NewDecoder(r io.Reader) runtime.Decoder {
	return &DecoderLocal{
		JSONPbDecoder: j.JSONPb.NewDecoder(r),
		JSONBuiltinDecoder: j.JSONBuiltin.NewDecoder(r),
	}
}

func (j *JSONLocal) NewEncoder(w io.Writer) runtime.Encoder {
	return j.JSONPb.NewEncoder(w)
}

func (j *JSONLocal) ContentType() string {
	return j.JSONPb.ContentType()
}

type DecoderLocal struct {
	JSONPbDecoder runtime.Decoder
	JSONBuiltinDecoder runtime.Decoder
}

func (d *DecoderLocal) Decode(v interface{}) error {
	if err := d.JSONBuiltinDecoder.Decode(v); err != nil && strings.Contains(err.Error(), `json: cannot unmarshal`) {
		return err
	}

	return d.JSONPbDecoder.Decode(v)
}

func (g *GrpcServer) Run() error {
	var attrs []zap.Field

	if g.name != `` {
		attrs = append(attrs, zap.String(`name`, g.name))
	}

	attrs = append(attrs,
		zap.Stringer(`addr`, g.listener.Addr()),
		zap.String(`listener`, reflect.TypeOf(g.listener).String()))

	for k, v := range g.attrs {
		attrs = append(attrs, zap.String(k, v))
	}

	g.logger.Info(`grpc run`, attrs...)

	// todo: poll health service
	close(g.ready)

	err := g.Server.Serve(g.listener)
	if err == grpc.ErrServerStopped {
		return nil
	}

	return err
}

func (g *GrpcServer) Close() error {
	g.logger.Info(`grpc gracefully stopping....`, zap.String(`addr`, g.addr))

	stopped := make(chan struct{})
	go func() {
		g.Server.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(g.GracefulStopTimeout)
	select {
	case <-t.C:
		g.logger.Info(`grpc ungracefully stopping....`, zap.String(`addr`, g.addr))
		g.Server.Stop()
	case <-stopped:
		t.Stop()
	}

	g.logger.Info(`grpc stopped`, zap.String(`addr`, g.addr))
	return nil
}

func (g *GrpcServer) Ready() <-chan struct{} {
	return g.ready
}

func (g *GrpcServer) LocalConnection() (*grpc.ClientConn, error) {
	if g.localConnection != nil {
		return g.localConnection, nil
	}

	var err error
	g.localConnection, err = grpc.Dial(
		g.listener.Addr().String(),
		grpc.WithInsecure(),
	)

	return g.localConnection, err
}

func (g *GrpcServer) EnableChannelZ() {
	channelzsvc.RegisterChannelzServiceToServer(g.Server)
}
