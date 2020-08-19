package usecase

import (
	"context"
	"time"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/domain/repo"
)

type usecaseImpl struct {
	TestTableRepo repo.TestTableRepo
}

func NewUsecase(testTableRepo repo.TestTableRepo) Usecase {
	return usecaseImpl{
		TestTableRepo: testTableRepo,
	}
}

func (u usecaseImpl) GetAllData(ctx context.Context) common.BaseResponse {

	time.Sleep(3 * time.Second)

	testTables, err := u.TestTableRepo.GetAll(ctx)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    testTables,
		Error:   nil,
	}
}

func (u usecaseImpl) CreateData(ctx context.Context, request entity.TestTable) common.BaseResponse {

	err := u.TestTableRepo.Create(ctx, request)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    nil,
		Error:   nil,
	}
}
