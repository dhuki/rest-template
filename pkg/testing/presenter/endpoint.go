package presenter

import (
	"context"

	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllDataEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return usecase.GetAllData(ctx)
	}
}
