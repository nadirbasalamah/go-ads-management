package response

import (
	"go-ads-management/businesses/users"
	"go-ads-management/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	CompanyName string         `json:"company_name"`
	Address     string         `json:"address"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	Role        utils.Role     `json:"role"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		CompanyName: domain.CompanyName,
		Address:     domain.Address,
		Username:    domain.Username,
		Email:       domain.Email,
		Password:    domain.Password,
		Role:        domain.Role,
	}
}
