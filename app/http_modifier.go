package app

import (
	"context"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	XHTTPCode             = "x-http-code"
	XHTTPCodeError        = "x-http-code-error"
	GRPCMetadataXHTTPCode = "Grpc-Metadata-X-Http-Code"
)

func SetHTTPErrorCode(ctx context.Context, httpStatusCode int) {
	_ = grpc.SetHeader(ctx, metadata.Pairs(XHTTPCodeError, strconv.Itoa(httpStatusCode)))
}

func SetHTTPSuccessCode(ctx context.Context, httpStatusCode int) {
	_ = grpc.SetHeader(ctx, metadata.Pairs(XHTTPCode, strconv.Itoa(httpStatusCode)))
}

func HTTPResponseErrorModifier(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	errCopy := err

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, errCopy)
		return
	}

	newReader, berr := utilities.IOReaderFactory(r.Body)
	if berr != nil {
		return
	}

	d := marshaler.NewDecoder(newReader())
	_ = d

	// set http status code
	if vals := md.HeaderMD.Get(XHTTPCodeError); len(vals) > 0 {
		var code int
		code, err = strconv.Atoi(vals[0])
		if err != nil {
			runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, errCopy)
			return
		}

		delete(md.HeaderMD, XHTTPCodeError)
		w.WriteHeader(code)
	}

	runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, errCopy)
}

func HTTPResponseModifier(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get(XHTTPCode); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}

		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, XHTTPCode)
		delete(w.Header(), GRPCMetadataXHTTPCode)
		w.WriteHeader(code)
	}

	return nil
}
