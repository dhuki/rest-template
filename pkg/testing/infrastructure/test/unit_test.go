package test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhuki/rest-template/config"
	"github.com/dhuki/rest-template/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template/pkg/testing/infrastructure"
)

// this is testing with mock connection database by sqlmock
// this sqlmock only check if query is correct with actual query or not

// command to get coverprofile if test in different package
// go test -coverpkg ./... ./test -coverprofile ./test/coverage.out
// to show in html format
// go tool cover -html=cp.out

func TestUnitGetAll(t *testing.T) {
	// using this default mathcer so it will using regex
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()

	// expected query
	mock.ExpectQuery(`[SELECT (.+) FROM test_tables]`).
		// will output from query above
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description"}).
			AddRow("testing", "testing"))

	gormDB, err := config.NewTesting(db)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	repo := infrastructure.NewTestTableInfrastructure(gormDB)
	if _, err := repo.GetAll(context.TODO()); err != nil {
		t.Errorf("%s", err)
		return
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUnitGet(t *testing.T) {
	// using this QueryMatcherEqual to ignore regex and become native query for postgres in this case
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()

	// expected query
	mock.ExpectQuery(`SELECT * FROM "test_tables" WHERE "test_tables"."id" = $1`).
		WithArgs(1).
		// will output from query above
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description"}).
			AddRow("testing", "testing"))

	gormDB, err := config.NewTesting(db)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	repo := infrastructure.NewTestTableInfrastructure(gormDB)
	if _, err := repo.Get(context.TODO()); err != nil {
		t.Errorf("%s", err)
		return
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUnitGetByName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT * FROM "test_tables" WHERE name = $1 ORDER BY "test_tables"."id" LIMIT 1`).
		WithArgs("testing").
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description"}).
			AddRow("testing", "testing"))

	gormDB, err := config.NewTesting(db)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	repo := infrastructure.NewTestTableInfrastructure(gormDB)
	if _, err := repo.GetByName(context.TODO(), "testing"); err != nil {
		t.Errorf("%s", err)
		return
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUnitCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer db.Close()

	mock.ExpectQuery(`INSERT INTO "test_tables" (.+) RETURNING`).
		WithArgs("testing", "testing").
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	gormDB, err := config.NewTesting(db)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	repo := infrastructure.NewTestTableInfrastructure(gormDB)
	if err := repo.Create(context.TODO(), entity.TestTable{
		Name:        "testing",
		Description: "testing",
	}); err != nil {
		t.Errorf("%s", err)
		return
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
