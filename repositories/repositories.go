package repositories

import "go-ads-management/models"

type UserRepository interface {
	Register(userInput models.RegisterInput) (models.User, error)
	GetByEmail(userInput models.LoginInput) (models.User, error)
	GetUserInfo(id string) (models.User, error)
}

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(categoryInput models.CategoryInput) (models.Category, error)
	Update(categoryInput models.CategoryInput, id string) (models.Category, error)
	Delete(id string) error
}

type AdsRepository interface {
	GetAll() ([]models.Ads, error)
	GetByID(id string) (models.Ads, error)
	Create(adsInput models.AdsInput) (models.Ads, error)
	Update(adsInput models.AdsInput, id string) (models.Ads, error)
	Delete(id string) error
	Restore(id string) (models.Ads, error)
	ForceDelete(id string) error
}
