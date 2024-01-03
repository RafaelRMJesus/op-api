package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDb() error {
	return errors.New("fake error")
}
