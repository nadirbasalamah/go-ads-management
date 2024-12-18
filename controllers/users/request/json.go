package request

import (
	"go-ads-management/businesses/users"
)

type UserRegister struct {
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		CompanyName: req.CompanyName,
		Address:     req.Address,
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
	}
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
