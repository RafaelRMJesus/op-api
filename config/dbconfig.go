package config

import (
	"os"

	"github.com/RafaelRMJesus/op-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	logger := GetLogger("dbconfig")
	dbPath := "./db/openings.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite_opening_error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("aqlite_automigration_error: %v", err)
		return nil, err
	}
	return db, nil
}
