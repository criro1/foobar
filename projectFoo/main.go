package main

import (
	"github.com/s7techlab/cckit/gateway"

	"github.com/criro1/foobar/app"
	"github.com/criro1/foobar/foo"
)

func main() {
	myApp, err := app.New("foo")
	if err != nil {
		panic(err)
	}

	fooServer := foo.NewServer()
	_, err = myApp.PublishServicesAndApplyConfig(app.AppConfig{Listen: app.Listen{HTTP: ":8081", GRPC: ":8082"}}, []gateway.ServiceDef{fooServer.ServiceDef()})
	if err != nil {
		panic(err)
	}

	myApp.MustServe()
}
