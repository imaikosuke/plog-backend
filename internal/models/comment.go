package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	PhotologID uint           `gorm:"not null" json:"photolog_id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	Content    string         `gorm:"not null" json:"content"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
