package entity

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID         uint           `gorm:"primary_key" json:"id"`
	Image      string         `json:"image"`
	CategoryId uint           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
