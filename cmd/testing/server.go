package testing

import (
	"github.com/dhuki/rest-template/pkg/testing/infrastructure"
	"github.com/dhuki/rest-template/pkg/testing/presenter"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/dhuki/rest-template/utils"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server interface {
	Start()
}

type testingServer struct {
	mux         *mux.Router
	db          *gorm.DB
	email       utils.Email
	middlewares []mux.MiddlewareFunc
	logger      log.Logger
}

func NewServer(mux *mux.Router, db *gorm.DB, email utils.Email, logger log.Logger, middlewares []mux.MiddlewareFunc) server {
	return testingServer{
		mux:         mux,
		db:          db,
		email:       email,
		logger:      logger,
		middlewares: middlewares,
	}
}

func (t testingServer) Start() {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewTestTableInfrastructure(t.db)
		srv = usecase.NewUsecase(infrastructure, t.email)
	}

	presenter.NewServer(t.mux, srv, t.logger, t.middlewares)
}
