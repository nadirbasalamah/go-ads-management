package models

import "github.com/go-playground/validator/v10"

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *LoginInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
