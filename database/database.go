package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initDB init database
func initDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
