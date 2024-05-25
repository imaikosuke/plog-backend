package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
    ID         uint           `gorm:"primaryKey" json:"id"`
    PhotologID uint           `gorm:"not null" json:"photolog_id"`
    ImageURL   string         `gorm:"not null" json:"image_url"`
    CreatedAt  time.Time      `json:"created_at"`
    UpdatedAt  time.Time      `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
