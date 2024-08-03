package services

import (
	"go-ads-management/models"
	"go-ads-management/repositories"

	"gorm.io/gorm"
)

type AdsService struct {
	repository repositories.AdsRepository
}

func InitAdsService() AdsService {
	return AdsService{
		repository: &repositories.AdsRepositoryImpl{},
	}
}

func (cs *AdsService) GetAll() (*gorm.DB, error) {
	return cs.repository.GetAll()
}

func (cs *AdsService) GetByID(id string) (models.Ads, error) {
	return cs.repository.GetByID(id)
}

func (cs *AdsService) Create(adsInput models.AdsInput) (models.Ads, error) {
	return cs.repository.Create(adsInput)
}

func (cs *AdsService) Update(adsInput models.AdsInput, id string) (models.Ads, error) {
	return cs.repository.Update(adsInput, id)
}

func (cs *AdsService) Delete(id string) error {
	return cs.repository.Delete(id)
}

func (cs *AdsService) Restore(id string) (models.Ads, error) {
	return cs.repository.Restore(id)
}

func (cs *AdsService) ForceDelete(id string) error {
	return cs.repository.ForceDelete(id)
}
