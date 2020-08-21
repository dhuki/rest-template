package usecase

import (
	"context"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/domain/repo"
	"github.com/dhuki/rest-template/utils"
)

// in usecaseImpl struct we just inject functionality of dependecy (dependent to abstraction (DI (Dependency Inversion)))
// not directly dependency itself (not dependent to lower module)
type usecaseImpl struct {
	TestTableRepo repo.TestTableRepo
	Email         utils.Email
}

func NewUsecase(testTableRepo repo.TestTableRepo, Email utils.Email) Usecase {
	return usecaseImpl{
		TestTableRepo: testTableRepo,
		Email:         Email,
	}
}

func (u usecaseImpl) GetAllData(ctx context.Context) common.BaseResponse {
	var response common.BaseResponse
	{
		// time.Sleep(3 * time.Second)
		testTables, err := u.TestTableRepo.GetAll(ctx)
		if err != nil {
			return common.BaseResponse{
				Error: err,
			}
		}

		response.Data = testTables
	}
	response.Success = common.RESPONSE_SUCCESS
	response.Message = common.RESPONSE_MSG_SUCCESS

	return response
}

func (u usecaseImpl) CreateData(ctx context.Context, request entity.TestTable) common.BaseResponse {

	err := u.TestTableRepo.Create(ctx, request)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	// try to send email
	go func() {
		u.Email.SendEmail(ctx, request)
	}()

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    nil,
		Error:   nil,
	}
}
