package response

import (
	"go-ads-management/businesses/ads"
	"time"

	"gorm.io/gorm"
)

type Ads struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	StartDate    string         `json:"start_date"`
	EndDate      string         `json:"end_date"`
	CategoryID   uint           `json:"category_id"`
	CategoryName string         `json:"category_name"`
	UserID       uint           `json:"user_id"`
	UserName     string         `json:"username"`
}

func FromDomain(domain ads.Domain) Ads {
	return Ads{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Title:        domain.Title,
		Description:  domain.Description,
		StartDate:    domain.StartDate,
		EndDate:      domain.EndDate,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		UserID:       domain.UserID,
		UserName:     domain.UserName,
	}
}
