package models

import (
	"time"

	"gorm.io/gorm"
)

type Ads struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Category    Category       `json:"category"`
	CategoryID  uint           `json:"category_id"`
	User        User           `json:"user"`
	UserID      uint           `json:"user_id"`
}
