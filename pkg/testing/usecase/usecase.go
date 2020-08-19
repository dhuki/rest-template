package usecase

import (
	"context"

	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
)

type Usecase interface {
	GetAllData(context.Context) common.BaseResponse
	CreateData(context.Context, entity.TestTable) common.BaseResponse
}
