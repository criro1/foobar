package app

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// HTTPHandlerMiddleware logs http requests and responses
func HTTPHandlerMiddleware(handler http.Handler, logger *zap.Logger, logRequestBody, logResponseBody bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fields := []zap.Field{
			zap.String(`method`, r.Method),
			zap.String(`url`, r.URL.String()),
		}

		requestFields := fields
		if logRequestBody {
			bb, _ := ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewReader(bb))
			if len(bb) > 0 {
				requestFields = append(fields, zap.String(`body`, string(bb)))
			}
		}

		logger.Debug(`http request started`, requestFields...)

		respWriter := &HttpResponseRecorder{ResponseWriter: w}
		handler.ServeHTTP(respWriter, r)

		fields = append(fields, zap.Int(`code`, respWriter.Status))
		if logResponseBody && r.Method != http.MethodGet {
			fields = append(fields, zap.String(`body`, string(respWriter.Body)))
		}

		if respWriter.Status >= 400 {
			logger.Error(`http request failed`, fields...)
		} else {
			logger.Info(`http request succeed`, fields...)
		}

	})
}

type HttpResponseRecorder struct {
	http.ResponseWriter
	Status      int
	Body        []byte
	wroteHeader bool
}

func (rr *HttpResponseRecorder) Write(bb []byte) (int, error) {
	rr.Body = bb
	rr.WriteHeader(http.StatusOK)

	return rr.ResponseWriter.Write(bb)
}

func (rr *HttpResponseRecorder) WriteHeader(code int) {
	if rr.wroteHeader {
		return
	}
	rr.wroteHeader = true
	rr.ResponseWriter.WriteHeader(code)
	rr.Status = code
}
