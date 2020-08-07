package config

import (
	"fmt"

	"github.com/dhuki/rest-template/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // side effect only run init function
)

func NewDatabase() (*gorm.DB, error) {
	// by default sslmode=enable, so you have to connect with ssl
	// since your server doesn't provide it
	// just use sslmode=disable
	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", common.DbHost, common.DbPort, common.DbName, common.DbUsername, common.DbPassword)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	// disable auto logging from gorm lib
	db.LogMode(false)

	return db, nil
}
