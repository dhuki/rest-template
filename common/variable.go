package common

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// create type of contextKey for key context
type (
	contextKey uint
)

var (
	DbUsername string
	DbPassword string
	DbName     string
	DbPort     string
	DbHost     string
)

// cons url
var (
	Host    string
	Port    string
	BaseUrl string
)

var (
	ErrDataNotFound = errors.New("Data not found")
	ErrAssertion    = errors.New("Error Assertion")
)

func LoadCons(path string) error {
	// by default you can use env directly by using "go env" to see list available value
	// but if you want use another key you should load another env file using
	// library "github.com/joho/godotenv".
	// if file env is in current dir just use .Load() only w/o parameter
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	Host = os.Getenv("base.host")
	Port = os.Getenv("base.port")
	BaseUrl = os.Getenv("base.url")

	DbUsername = os.Getenv("db.username")
	DbPassword = os.Getenv("db.password")
	DbName = os.Getenv("db.name")
	DbPort = os.Getenv("db.port")
	DbHost = os.Getenv("db.host")

	return nil
}
