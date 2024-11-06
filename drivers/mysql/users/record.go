package users

import (
	"go-ads-management/businesses/users"
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
	Username    string         `json:"username"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"-"`
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:          rec.ID,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
		CompanyName: rec.CompanyName,
		Address:     rec.Address,
		Username:    rec.Username,
		Email:       rec.Email,
		Password:    rec.Password,
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		CompanyName: domain.CompanyName,
		Address:     domain.Address,
		Username:    domain.Username,
		Email:       domain.Email,
		Password:    domain.Password,
	}
}