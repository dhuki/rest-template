package usecase

import (
	"context"

	"github.com/dhuki/rest-template/common"
)

type Usecase interface {
	GetAllData(context.Context) common.BaseResponse
}
