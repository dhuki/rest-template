package testing

import (
	"github.com/dhuki/rest-template/pkg/testing/infrastructure"
	"github.com/dhuki/rest-template/pkg/testing/presenter"
	"github.com/dhuki/rest-template/pkg/testing/usecase"
	"github.com/dhuki/rest-template/utils"
	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server interface {
	Start()
}

type testingServer struct {
	mux         *mux.Router
	db          *gorm.DB
	redisClient *redis.Client
	middlewares []mux.MiddlewareFunc
	utils       *utils.Utils
	logger      log.Logger
}

func NewServer(mux *mux.Router, redisClient *redis.Client, logger log.Logger) testingServer {
	return testingServer{
		mux:         mux,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (t testingServer) AddDatabase(db *gorm.DB) testingServer {
	t.db = db
	return t
}

func (t testingServer) AddUtils(utils *utils.Utils) testingServer {
	t.utils = utils
	return t
}

func (t testingServer) AddMiddlewares(middlewares []mux.MiddlewareFunc) testingServer {
	t.middlewares = middlewares
	return t
}

func (t testingServer) Build() server {
	return t
}

func (t testingServer) Start() {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewTestTableInfrastructure(t.db)
		srv = usecase.NewUsecase(infrastructure, t.utils)
	}

	presenter.NewServer(t.mux, srv, t.logger, t.middlewares)
}
