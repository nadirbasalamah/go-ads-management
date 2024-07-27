package models

import "github.com/go-playground/validator/v10"

type RegisterInput struct {
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

func (u *RegisterInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
