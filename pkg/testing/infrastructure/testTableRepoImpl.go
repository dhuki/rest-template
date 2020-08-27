package infrastructure

import (
	"context"

	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/domain/repo"
	"gorm.io/gorm"
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
	// db := t.db.Find(&testTables) // this is ver1 of gorm cannot use context
	db := t.db.WithContext(ctx).Find(&testTables) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return nil, db.Error
	}

	return testTables, nil
}

func (t testTableRepoImpl) Get(ctx context.Context) (entity.TestTable, error) {
	testTables := entity.TestTable{
		ID: 1,
	}
	// db := t.db.Find(&testTables) // this is ver1 of gorm cannot use context
	db := t.db.WithContext(ctx).Find(&testTables) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return testTables, db.Error
	}

	return testTables, nil
}

func (t testTableRepoImpl) GetByName(ctx context.Context, name string) (entity.TestTable, error) {
	var testTables entity.TestTable
	// db := t.db.Find(&testTables) // this is ver1 of gorm cannot use context
	db := t.db.WithContext(ctx).Where("name = ?", name).First(&testTables) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return testTables, db.Error
	}

	return testTables, nil
}

func (t testTableRepoImpl) Create(ctx context.Context, testTable entity.TestTable) error {
	db := t.db.WithContext(ctx).Create(&testTable) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return db.Error
	}
	return nil
}
