package request

import (
	"go-ads-management/businesses/categories"
)

type Category struct {
	Name string `json:"name" validate:"required"`
}

func (req *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name: req.Name,
	}
}
