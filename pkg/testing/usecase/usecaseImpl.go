package usecase

import (
	"context"

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

func (u usecaseImpl) GetAllData(ctx context.Context) (common.BaseResponse, error) {
	var response common.BaseResponse
	{
		testTables, err := u.TestTableRepo.GetAll(ctx)
		if err != nil {
			return response, err
		}

		response.Data = testTables
	}
	response.Success = common.RESPONSE_SUCCESS
	response.Message = common.RESPONSE_MSG_SUCCESS

	return response, nil
}
