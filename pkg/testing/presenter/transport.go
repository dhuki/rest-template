package presenter

import (
	"github.com/dhuki/rest-template/middleware"
	"github.com/dhuki/rest-template/pkg/testing/presenter/model"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewServer(mux *mux.Router, usecase usecase.Usecase, logger log.Logger) {
	r := mux.PathPrefix("/api").Subrouter()
	r.Use(middleware.SetContentTypeToJson)
	r.Use(middleware.SetInterceptors(logger))

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(middleware.ErrorEncoder),             // error for client
		httptransport.ServerErrorHandler(middleware.ErrorHandlerFunc(logger)), // error for internal
	}

	r.Methods("GET").Path("/testing").Handler(httptransport.NewServer(
		MakeGetAllDataEndpoint(usecase),
		model.DecodeGetAllRequest,
		model.EncodeResponse,
		options...,
	))

	r.Methods("GET").Path("/testing/{param}").Handler(httptransport.NewServer(
		MakeGetAllDataEndpoint(usecase),
		model.DecodeGetAllRequest,
		model.EncodeResponse,
		options...,
	))

	r.Methods("GET").Path("/testing").Queries(
		"param", "{param}",
	).Handler(httptransport.NewServer(
		MakeGetAllDataEndpoint(usecase),
		model.DecodeGetAllRequest,
		model.EncodeResponse,
		options...,
	))
}
