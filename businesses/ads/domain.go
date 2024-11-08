package ads

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title        string
	Description  string
	StartDate    string
	EndDate      string
	CategoryID   uint
	CategoryName string
	UserID       uint
	UserName     string
	MediaURL     string
	MediaCID     string
	MediaID      string
}

type UseCase interface {
	GetAll(ctx context.Context) (*gorm.DB, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByCategory(ctx context.Context, categoryID string) (*gorm.DB, error)
	GetByUser(ctx context.Context) (*gorm.DB, error)
	GetTrashed(ctx context.Context) (*gorm.DB, error)
	Create(ctx context.Context, adsReq *Domain) (Domain, error)
	Update(ctx context.Context, adsReq *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) (*gorm.DB, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByCategory(ctx context.Context, categoryID string) (*gorm.DB, error)
	GetByUser(ctx context.Context) (*gorm.DB, error)
	GetTrashed(ctx context.Context) (*gorm.DB, error)
	Create(ctx context.Context, adsReq *Domain) (Domain, error)
	Update(ctx context.Context, adsReq *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}
