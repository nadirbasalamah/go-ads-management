package ads

import (
	"go-ads-management/businesses/ads"
	"go-ads-management/drivers/mysql/categories"
	"go-ads-management/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Ads struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	StartDate   string              `json:"start_date"`
	EndDate     string              `json:"end_date"`
	Category    categories.Category `json:"category"`
	CategoryID  uint                `json:"category_id"`
	User        users.User          `json:"user"`
	UserID      uint                `json:"user_id"`
	MediaURL    string              `json:"media_url"`
	MediaCID    string              `json:"media_cid" gorm:"column:media_cid"`
	MediaID     string              `json:"media_id"`
}

func (rec *Ads) ToDomain() ads.Domain {
	return ads.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		Title:        rec.Title,
		Description:  rec.Description,
		StartDate:    rec.StartDate,
		EndDate:      rec.EndDate,
		CategoryID:   rec.Category.ID,
		CategoryName: rec.Category.Name,
		UserID:       rec.User.ID,
		UserName:     rec.User.Username,
		MediaURL:     rec.MediaURL,
		MediaCID:     rec.MediaCID,
		MediaID:      rec.MediaID,
	}
}

func FromDomain(domain *ads.Domain) *Ads {
	return &Ads{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		StartDate:   domain.StartDate,
		EndDate:     domain.EndDate,
		CategoryID:  domain.CategoryID,
		UserID:      domain.UserID,
		MediaURL:    domain.MediaURL,
		MediaCID:    domain.MediaCID,
		MediaID:     domain.MediaID,
	}
}
