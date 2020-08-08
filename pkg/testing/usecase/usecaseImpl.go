package usecase

import (
	"context"
	"time"

	"github.com/dhuki/rest-template/common"
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

	testTables, err := u.TestTableRepo.GetAll(ctx)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	time.Sleep(time.Second * 2)

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    testTables,
		Error:   nil,
	}
}
