package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dhuki/rest-template/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
)

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
