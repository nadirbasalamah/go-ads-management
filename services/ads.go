package services

import (
	"go-ads-management/models"
	"go-ads-management/repositories"

	"gorm.io/gorm"
)

type AdsService struct {
	Repository repositories.AdsRepository
}

func InitAdsService() AdsService {
	return AdsService{
		Repository: &repositories.AdsRepositoryImpl{},
	}
}

func (cs *AdsService) GetAll() (*gorm.DB, error) {
	return cs.Repository.GetAll()
}

func (cs *AdsService) GetByID(id string) (models.Ads, error) {
	return cs.Repository.GetByID(id)
}

func (cs *AdsService) Create(adsInput models.AdsInput) (models.Ads, error) {
	return cs.Repository.Create(adsInput)
}

func (cs *AdsService) Update(adsInput models.AdsInput, id string) (models.Ads, error) {
	return cs.Repository.Update(adsInput, id)
}

func (cs *AdsService) Delete(id string) error {
	return cs.Repository.Delete(id)
}

func (cs *AdsService) Restore(id string) (models.Ads, error) {
	return cs.Repository.Restore(id)
}

func (cs *AdsService) ForceDelete(id string) error {
	return cs.Repository.ForceDelete(id)
}
