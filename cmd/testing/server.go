package testing

import (
	"github.com/dhuki/rest-template/pkg/testing/infrastructure"
	"github.com/dhuki/rest-template/pkg/testing/presenter"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type server interface {
	Start()
}

type testingServer struct {
	mux    *mux.Router
	db     *gorm.DB
	logger log.Logger
}

func NewServer(mux *mux.Router, db *gorm.DB, logger log.Logger) server {
	return testingServer{
		mux:    mux,
		db:     db,
		logger: logger,
	}
}

func (t testingServer) Start() {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewTestTableInfrastructure(t.db)
		srv = usecase.NewUsecase(infrastructure)
	}

	presenter.NewServer(t.mux, srv, t.logger)
}
