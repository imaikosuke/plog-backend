package models

import (
	"time"

	"github.com/lib/pq"
)

type Photolog struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"primaryKey;not null" json:"user_id"`
	GeneratedText string         `gorm:"not null" json:"generated_text"`
	Images        pq.StringArray `gorm:"type:text[]" json:"images"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
