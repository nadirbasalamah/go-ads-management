package categories

import (
	"errors"
	"go-ads-management/businesses/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
}

func (rec *Category) BeforeDelete(tx *gorm.DB) error {
	var count int64

	err := tx.Table("ads").Where("category_id = ?", rec.ID).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("cannot delete category: it is still being used in advertisements")
	}

	return nil
}

func (rec *Category) ToDomain() categories.Domain {
	return categories.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
	}
}

func FromDomain(domain *categories.Domain) *Category {
	return &Category{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
	}
}
