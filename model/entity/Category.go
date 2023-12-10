package entity

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Name      string         `json:"name"`
	Photos    []Photo        `json:"photos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
