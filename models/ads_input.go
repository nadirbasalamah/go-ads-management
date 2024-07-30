package models

import "github.com/go-playground/validator/v10"

type AdsInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	UserID      uint   `json:"user_id"`
}

func (a *AdsInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(a)

	return err
}
