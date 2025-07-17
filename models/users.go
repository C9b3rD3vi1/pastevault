package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
