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
	"github.com/go-kit/kit/log/level"
	"gorm.io/gorm"
)

func main() {
	errChan := make(chan error, 1) // buffered channel
	dbChan := make(chan *gorm.DB)

	// set up logger
	logger := config.NewLogger()
	level.Info(logger).Log("description", "Server is running")
	defer level.Info(logger).Log("description", "Server shutdown")

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

	// set up router configuration
	go func() {
		db := <-dbChan
		router := config.NewRouter()

		testing.NewServer(router.Mux, db, logger).Start()

		level.Info(logger).Log("description", fmt.Sprintf("Listen to port :%s", common.Port))

		errChan <- router.Start()
	}()

	level.Error(logger).Log("description", "Server error", "message", <-errChan)

	fmt.Println(runtime.NumGoroutine())
}
