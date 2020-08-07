package infrastructure

import (
	"context"

	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/domain/repo"
	"github.com/jinzhu/gorm"
)

// package that implement repo in domain layer
// it places outermost (paling luar) of all layer

type testTableRepoImpl struct {
	db *gorm.DB
}

func NewTestTableInfrastructure(db *gorm.DB) repo.TestTableRepo {
	return testTableRepoImpl{
		db: db,
	}
}

func (t testTableRepoImpl) GetAll(ctx context.Context) ([]entity.TestTable, error) {
	var testTables []entity.TestTable
	db := t.db.Find(&testTables)
	if db.Error != nil {
		return nil, db.Error
	}

	return testTables, nil
}
