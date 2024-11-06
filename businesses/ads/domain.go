package ads

import (
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
}

type UseCase interface {
	GetAll() (*gorm.DB, error)
	GetByID(id string) (Domain, error)
	Create(adsReq *Domain) (Domain, error)
	Update(adsReq *Domain, id string) (Domain, error)
	Delete(id string) error
	Restore(id string) (Domain, error)
	ForceDelete(id string) error
}

type Repository interface {
	GetAll() (*gorm.DB, error)
	GetByID(id string) (Domain, error)
	Create(adsReq *Domain) (Domain, error)
	Update(adsReq *Domain, id string) (Domain, error)
	Delete(id string) error
	Restore(id string) (Domain, error)
	ForceDelete(id string) error
}
