package models

import (
	"time"

	"gorm.io/gorm"
)

type Photolog struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    UserID        uint           `gorm:"not null" json:"user_id"`
    GeneratedText string         `gorm:"not null" json:"generated_text"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
