package request

import (
	"go-ads-management/businesses/ads"

	"github.com/go-playground/validator/v10"
)

type Ads struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	UserID      uint   `json:"user_id"`
}

func (req *Ads) ToDomain() *ads.Domain {
	return &ads.Domain{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CategoryID:  req.CategoryID,
		UserID:      req.UserID,
	}
}

func (req *Ads) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
