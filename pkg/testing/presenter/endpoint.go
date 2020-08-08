package presenter

import (
	"context"
	"fmt"
	"runtime"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllDataEndpointWithGoroutine(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
		fmt.Println(runtime.NumGoroutine())
		return response, response.Error
	}
}
