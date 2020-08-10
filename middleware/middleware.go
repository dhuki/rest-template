package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dhuki/rest-template/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
)

// common middleware return type as json
func SetContentTypeToJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // set up content-type of json at header for response
		next.ServeHTTP(w, r)
	})
}

// common middleware return type as json
func SetInterceptors(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func(begin time.Time) {
				level.Info(logger).Log(
					"description", "INTERCEPTOR",
					"scheme", func() string {
						if r.TLS == nil {
							return "http"
						}
						return "https"
					}(),
					"took", time.Since(begin),
					"host", r.Host,
					"url", r.URL.String(),
					"method", r.Method,
					"requestBody", fmt.Sprintf("%+v", r.Body)) // givin output of struct to this -> attribute : value
				// "response", fmt.Sprintf("%+v", w.)) // givin output of struct to this -> attribute : value
			}(time.Now())

			next.ServeHTTP(w, r)
		})
	}
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	var response common.BaseResponse
	{
		switch err {
		case common.ErrDataNotFound:
			response.Data = common.ErrDataNotFound.Error()
		}
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorHandlerFunc(logger log.Logger) transport.ErrorHandlerFunc {
	return func(_ context.Context, err error) {
		level.Error(logger).Log(
			"description", "Internal server error",
			"message", err,
			"solution", "Please check encode/decode body, usecase method, and dependency library",
		)
	}
}
