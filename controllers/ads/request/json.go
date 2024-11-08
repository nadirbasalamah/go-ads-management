package request

import (
	"go-ads-management/businesses/ads"
	"mime/multipart"
)

type Ads struct {
	Title       string `form:"title" validate:"required"`
	Description string `form:"description" validate:"required"`
	StartDate   string `form:"start_date" validate:"required,validDate"`
	EndDate     string `form:"end_date" validate:"required,validDate"`
	CategoryID  uint   `form:"category_id" validate:"required"`
	File        *multipart.FileHeader
	MediaURL    string
	MediaCID    string
}

func (req *Ads) ToDomain() *ads.Domain {
	return &ads.Domain{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CategoryID:  req.CategoryID,
		MediaURL:    req.MediaURL,
		MediaCID:    req.MediaCID,
	}
}
