package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/dhuki/rest-template/cmd/testing"
	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/config"
	"github.com/dhuki/rest-template/middleware"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"
	"gorm.io/gorm"
)

func main() {
	errChan := make(chan error, 1) // buffered channel
	dbChan := make(chan *gorm.DB)
	bucketChan := make(chan *ratelimit.Bucket)

	// set up logger
	logger := config.NewLogger()
	level.Info(logger).Log("description", "Server is running")
	defer level.Info(logger).Log("description", "Server shutdown")

	// load common cons from .env
	err := common.LoadCons(common.ENV_PATH)
	if err != nil {
		errChan <- err
	}

	// setup command kill and interrupt
	go func() {
		sign := make(chan os.Signal) // buffered channel
		// SIGINT (Signal Interrupt (CTRL + C))
		// SIGTERM (Signal Terminated (KILL command))
		signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-sign)
	}()

	// set up database configuration
	go func() {
		db, err := config.NewDatabase()
		if err != nil {
			errChan <- err
		}
		dbChan <- db
	}()

	// set up rate limiter configuration
	go func() {
		bucketChan <- config.NewRateLimit()
	}()

	// set up router configuration
	go func() {
		db := <-dbChan
		bucketToken := <-bucketChan

		router := config.NewRouter()

		// set up module with dependencies
		testing.NewServer(router.Mux, db, logger, []mux.MiddlewareFunc{ // list of needed middleware, order is matter
			handlers.CompressHandler,
			middleware.TokenBucketLimiter(bucketToken, logger),
		}).Start()

		level.Info(logger).Log("description", fmt.Sprintf("Listen to port :%s", common.Port))

		errChan <- router.Start()
	}()

	level.Error(logger).Log("description", "Server error", "message", <-errChan)

	fmt.Println(runtime.NumGoroutine())
}
