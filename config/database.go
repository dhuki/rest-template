package config

import (
	"fmt"

	"github.com/dhuki/rest-template/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	// by default sslmode=enable, so you have to connect with ssl
	// since your server doesn't provide it
	// just use sslmode=disable
	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", common.DbHost, common.DbPort, common.DbName, common.DbUsername, common.DbPassword)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // disable auto log gorm v2
	}) // configuration gorm v2 (it's unstable yet)
	if err != nil {
		return nil, err
	}
	// disable auto logging from gorm lib v1
	// db.LogMode(false)

	return db, nil
}
