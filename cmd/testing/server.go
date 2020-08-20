package testing

import (
	"github.com/dhuki/rest-template/pkg/testing/infrastructure"
	"github.com/dhuki/rest-template/pkg/testing/presenter"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
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
	middlewares []mux.MiddlewareFunc
	logger      log.Logger
}

func NewServer(mux *mux.Router, db *gorm.DB, logger log.Logger, middlewares []mux.MiddlewareFunc) server {
	return testingServer{
		mux:         mux,
		db:          db,
		logger:      logger,
		middlewares: middlewares,
	}
}

func (t testingServer) Start() {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewTestTableInfrastructure(t.db)
		srv = usecase.NewUsecase(infrastructure)
	}

	presenter.NewServer(t.mux, srv, t.logger, t.middlewares)
}
