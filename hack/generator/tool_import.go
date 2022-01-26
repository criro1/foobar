// +build tools

package proto_gen

import (
	// proto/grpc
	_ "github.com/golang/protobuf/protoc-gen-go"
	// json gateway
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// docs
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
	// validation schema
	_ "github.com/mwitkow/go-proto-validators/protoc-gen-govalidators"
	// chaincode staff
	_ "github.com/s7techlab/cckit/gateway/protoc-gen-cc-gateway"
	// cli staff
	_ "github.com/fiorix/protoc-gen-cobra"
	// packr
	_ "github.com/gobuffalo/packr/packr"

	_ "github.com/kazegusuri/channelzcli"
)
