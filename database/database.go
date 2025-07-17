package database

import (
	"log"

	"github.com/C9b3rD3vi1/pastevault/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initDB init database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/secrets.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		//return nil, err
	}
	// Set the global DB variable
	DB = db
	// Log the database connection
	log.Println("Connected to the database")

	// Auto-migrate your Secret model
	err = db.AutoMigrate(&models.Secret{}, &models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	// Log the migration
	log.Println("Database schema migrated successfully")

	return db, nil
}
