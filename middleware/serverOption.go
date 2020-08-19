package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dhuki/rest-template/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

func SetInterceptors(logger log.Logger) httptransport.ServerFinalizerFunc {
	return httptransport.ServerFinalizerFunc(func(ctx context.Context, code int, r *http.Request) {
		level.Info(logger).Log(
			"description", "Interceptors",
			"scheme", ctx.Value(httptransport.ContextKeyRequestProto),
			"host", ctx.Value(httptransport.ContextKeyRequestHost),
			"URI", ctx.Value(httptransport.ContextKeyRequestURI),
			"path", ctx.Value(httptransport.ContextKeyRequestPath),
			"method", ctx.Value(httptransport.ContextKeyRequestMethod),
			"statusCode", code)
		// "requestBody", fmt.Sprintf("%+v", request)) // givin output of struct to this -> attribute : value
		// "response", fmt.Sprintf("%+v", w.)) // givin output of struct to this -> attribute : value
	})
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	var response common.BaseResponse
	{
		switch err {
		case common.ErrDataNotFound:
			response.Data = common.ErrDataNotFound.Error()
		case common.ErrAssertion:
			response.Data = common.ErrAssertion.Error()
		}
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorHandlerFunc(logger log.Logger) transport.ErrorHandlerFunc {
	return transport.ErrorHandlerFunc(func(_ context.Context, err error) {
		level.Error(logger).Log(
			"description", "Internal server error",
			"message", err,
			"solution", "Please check encode/decode body, usecase method, and dependency library",
		)
	})
}
