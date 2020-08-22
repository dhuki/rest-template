package utils

import (
	"net/http"

	"gorm.io/gorm"
)

type Cleanup interface {
	Close() error
}

type Dependencies struct {
	GormDB *gorm.DB
	Server *http.Server
}

func (d Dependencies) Close() error {
	if err := d.Server.Close(); err != nil {
		return err
	}
	if _, err := d.GormDB.DB(); err != nil {
		return err
	}
	return nil
}
