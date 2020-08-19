package presenter

import (
	"context"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllDataEndpointWithGoroutine(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// source : https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
		response := make(chan common.BaseResponse, 1)

		go func() {
			response <- usecase.GetAllData(ctx)
		}()

		select {
		case <-ctx.Done():
			return nil, nil
		case result := <-response:
			if result.Error != nil {
				return nil, result.Error
			}
			return result, nil
		}
	}
}

func MakeGetAllDataEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response := usecase.GetAllData(ctx)
		return response, response.Error
	}
}

func MakeCreateDataEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(entity.TestTable)
		if !ok {
			return nil, common.ErrAssertion
		}
		response := usecase.CreateData(ctx, req)
		return response, response.Error
	}
}
