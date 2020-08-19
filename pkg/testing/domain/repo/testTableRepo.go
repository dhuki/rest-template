package repo

import (
	"context"

	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
)

// package for defined repository that able to use by entity

type TestTableRepo interface {
	GetAll(context.Context) ([]entity.TestTable, error)
	Create(context.Context, entity.TestTable) error
}
