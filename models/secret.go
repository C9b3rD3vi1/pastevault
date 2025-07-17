package models

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"uniqueIndex"`
	Password  string    `gorm:"uniqueIndex"`
	Content   string    `gorm:"type:text"`
	Viewed    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ExpiresAt time.Time `gorm:"autoCreateTime"`
}
