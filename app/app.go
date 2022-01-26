package app

import (
	"fmt"
	"os"

	"github.com/s7techlab/cckit/gateway"
	"go.uber.org/zap"
)

type (
	App struct {
		name string

		Http []*HttpServer
		Grpc []*GrpcServer

		Workers *Pool
		logger  *zap.Logger
	}

	AppOpt func(*App) error
)

func MustNew(name string, opts ...AppOpt) *App {
	app, err := New(name, opts...)
	if err != nil {
		panic(err)
	}

	return app
}

func New(name string, opts ...AppOpt) (*App, error) {
	app := &App{
		name: name,
	}
	if app.logger == nil {
		app.logger = L()
	}
	app.Workers = NewPool(WithLogger(app.logger))

	for _, opt := range opts {
		if err := opt(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *App) Logger() *zap.Logger {
	return a.logger
}

func (a *App) Run() error {
	a.logger.Info(`app starting`,
		zap.Int(`runners`, len(a.Workers.Runners())),
		zap.Int(`closers`, len(a.Workers.Closers())))

	a.Workers.ShutdownOnSignal()

	if err := a.Workers.Run(); err != nil {
		return err
	}

	a.logger.Info(`app started`)

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) MustServe() {
	a.MustRun()

	<-a.Workers.Ready()
	a.logger.Info(`runners ready`)

	os.Exit(<-a.Workers.Done())
}

func WithPublishService(listen Listen, serviceDefs []gateway.ServiceDef, opts ...PublisherOpt) func(*App) error {
	return func(app *App) error {
		_, err := PublishServices(app, listen, serviceDefs, opts...)
		return err
	}
}

func PublishServices(app *App, listen Listen, serviceDefs []gateway.ServiceDef, opts ...PublisherOpt) (*ServicePublisher, error) {
	publisher, err := NewServicePublisher(listen, serviceDefs, app.name, app.logger, opts...)
	if err != nil {
		return nil, fmt.Errorf(`service publisher: %w`, err)
	}

	if publisher.Http != nil {
		app.AddHttp(publisher.Http)
	}

	if publisher.Grpc != nil {
		app.AddGrpc(publisher.Grpc)
	}

	return publisher, nil
}

func (a *App) PublishServices(listen Listen, serviceDefs []gateway.ServiceDef, opts ...PublisherOpt) (*ServicePublisher, error) {
	a.logger.Info(`publish services`, zap.Reflect(`address`, listen))
	return PublishServices(a, listen, serviceDefs, opts...)
}

func (a *App) PublishServicesAndApplyConfig(conf AppConfig, serviceDefs []gateway.ServiceDef, opts ...PublisherOpt) (*ServicePublisher, error) {
	publisher, err := a.PublishServices(conf.Listen, serviceDefs, opts...)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func (a *App) AddHttp(http *HttpServer) {
	a.Workers.Add(http)
	a.Http = append(a.Http, http)
}

func (a *App) AddGrpc(grpc *GrpcServer) {
	a.Workers.Add(grpc)
	a.Grpc = append(a.Grpc, grpc)
}

func (a *App) HandleHTTP(addr string, opts ...HTTPServerOpt) (*HttpServer, error) {
	http, err := NewHTTP(addr, a.logger, opts...)
	if err != nil {
		return nil, err
	}

	a.AddHttp(http)
	return http, nil
}

func (a *App) HandleGrpc(addr string, opts ...GRPCServerOpt) (*GrpcServer, error) {
	grpc, err := NewGrpc(addr, a.logger, opts...)
	if err != nil {
		return nil, err
	}

	a.AddGrpc(grpc)
	return grpc, nil
}
