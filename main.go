package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/dhuki/rest-template/cmd/testing"
	"github.com/dhuki/rest-template/common"
	"github.com/dhuki/rest-template/config"
	"github.com/dhuki/rest-template/middleware"
	"github.com/dhuki/rest-template/utils"
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
	emailChan := make(chan utils.Email)
	dependenciesChan := make(chan utils.Dependencies)

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
		fmt.Println("Close goroutine signal")
	}()

	// set up database configuration
	go func() {
		db, err := config.NewDatabase()
		if err != nil {
			errChan <- err
		}
		dbChan <- db
		fmt.Println("Close goroutine database")
	}()

	// set up rate limiter configuration
	go func() {
		bucketChan <- config.NewRateLimit()
		fmt.Println("Close goroutine rate limiter")
	}()

	// set up email configuration and email util
	go func() {
		smtpAuth := config.NewEmail()
		emailChan <- utils.NewEmailUtil(smtpAuth, logger)
		fmt.Println("Close goroutine email")
	}()

	go func() {
		db := <-dbChan
		email := <-emailChan
		bucketToken := <-bucketChan

		// set up router configuration
		router := config.NewRouter()
		// set up module with dependencies
		testing.NewServer(router.Mux, db, email, logger, []mux.MiddlewareFunc{
			// list of middleware that needed, order is matter
			handlers.CompressHandler,
			middleware.TokenBucketLimiter(bucketToken, logger),
		}).Start()

		dependenciesChan <- utils.Dependencies{
			GormDB: db,
			Server: router.Server,
		}

		errChan <- router.Start()
		fmt.Println("Close goroutine router")
	}()

	dependencies := <-dependenciesChan
	level.Info(logger).Log("description", fmt.Sprintf("Listen to port :%s", common.Port))
	level.Error(logger).Log("description", "Server error", "message", <-errChan)

	if err := dependencies.Close(); err != nil {
		level.Error(logger).Log("description", "Server Cannot Close Dependencies", "message", err)
	}

	time.Sleep(3 * time.Second)

	// print number of goroutine that running
	fmt.Println(runtime.NumGoroutine())
}
