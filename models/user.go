package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CompanyName string         `json:"company_name"`
	Address     string         `json:"address"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"-"`
}
