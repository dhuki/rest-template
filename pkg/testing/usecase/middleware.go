package usecase

import (
	"context"
	"time"

	"github.com/dhuki/rest-template/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type middelware struct {
	usecase Usecase
	logger  log.Logger
}

func NewMiddleware(usecase Usecase, logger log.Logger) Usecase {
	return middelware{
		usecase: usecase,
		logger:  logger,
	}
}

func (m middelware) GetAllData(ctx context.Context) (response common.BaseResponse) {
	// interceptors
	defer func(begin time.Time) {
		level.Info(m.logger).Log(
			"description", "Interceptors",
			"took", time.Since(begin),
		)
	}(time.Now())

	return m.usecase.GetAllData(ctx)
}
