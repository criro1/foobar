package foo

import (
	"context"
	"fmt"

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

	var (
		fir []string
		sec []int64
		th string
	)


	switch vt := req.BarRequests.(type) {
	case *BarRequest_First:
		fir = vt.First.Id
		if len(fir) == 0 {
			return nil, fmt.Errorf("first len of stings is 0")
		}

	case *BarRequest_Second:
		sec = vt.Second.Num
		if len(sec) == 0 {
			return nil, fmt.Errorf("second len of int64 is 0")
		}
	case *BarRequest_Third:
		th = vt.Third
		if th == "" {
			return nil, fmt.Errorf("third string is empty")
		}
	default:
		return nil, fmt.Errorf("no one type")
	}

	return &BarResponse{Text: "this is Bar"}, nil
}

func (s *FooServerLocal) ServiceDef() gateway.ServiceDef {
	return gateway.ServiceDef{
		Desc:                        &_Foo_serviceDesc,
		Service:                     s,
		HandlerFromEndpointRegister: RegisterFooHandlerFromEndpoint,
	}
}
