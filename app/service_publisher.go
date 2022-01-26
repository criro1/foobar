package app

import (
	"github.com/google/martian/cors"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/s7techlab/cckit/gateway"
	"go.uber.org/zap"
)

type ServicePublisher struct {
	Http *HttpServer
	Grpc *GrpcServer
}

type (
	publisherOpts struct {
		grpcOpts []GRPCServerOpt
	}
	PublisherOpt func(opts *publisherOpts) error
)

func WithGRPCServerOpt(opt GRPCServerOpt) PublisherOpt {
	return func(opts *publisherOpts) error {
		opts.grpcOpts = append(opts.grpcOpts, opt)

		return nil
	}
}

func NewServicePublisher(
	listen Listen,
	serviceDefs []gateway.ServiceDef,
	appName string,
	logger *zap.Logger,
	opts ...PublisherOpt,
) (*ServicePublisher, error) {
	var (
		publisherOptsInstance = new(publisherOpts)
		publisher             = new(ServicePublisher)
		serveMux              *runtime.ServeMux
		err                   error
	)
	for _, opt := range opts {
		if err = opt(publisherOptsInstance); err != nil {
			return nil, err
		}
	}
	publisher.Grpc, err = NewGrpc(listen.GRPC, logger, publisherOptsInstance.grpcOpts...)
	if err != nil {
		return nil, err
	}

	// http не выставляем
	if listen.HTTP == `` {
		publisher.Grpc.RegisterServices(serviceDefs)
		publisher.Http = NewDisabledHTTP(logger)
	} else {
		publisher.Http, err = NewHTTP(listen.HTTP, logger)
		if err != nil {
			return nil, err
		}

		serveMux, err = publisher.Grpc.RegisterServicesToMux(serviceDefs)
		if err != nil {
			return nil, err
		}

		publisher.Http.Mux.Handle(`/`, cors.NewHandler(serveMux))

		// todo: add swagger from service def
	}

	return publisher, nil
}
