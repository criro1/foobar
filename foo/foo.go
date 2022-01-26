package foo

import (
	"context"

	"github.com/s7techlab/cckit/gateway"
)

type Matcher func(ctx context.Context)

func NewServer() *FooServerLocal {
	return &FooServerLocal{
		matcher: func(ctx context.Context) {},
	}
}

type FooServerLocal struct {
	matcher Matcher
}

func (s *FooServerLocal) SetMatcher(matcher Matcher) {
	s.matcher = matcher
}

func (s *FooServerLocal) Bar(ctx context.Context, req *BarRequest) (*BarResponse, error) {
	s.matcher(ctx)

	return &BarResponse{Text: "this is Bar"}, nil
}

func (s *FooServerLocal) ServiceDef() gateway.ServiceDef {
	return gateway.ServiceDef{
		Desc:                        &_Foo_serviceDesc,
		Service:                     s,
		HandlerFromEndpointRegister: RegisterFooHandlerFromEndpoint,
	}
}
