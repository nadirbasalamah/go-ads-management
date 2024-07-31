package repositories

import (
	"go-ads-management/database"
	"go-ads-management/models"
)

type AdsRepositoryImpl struct {
}

func InitAdsRepository() AdsRepository {
	return &AdsRepositoryImpl{}
}

func (cr *AdsRepositoryImpl) GetAll() ([]models.Ads, error) {
	var ads []models.Ads

	if err := database.DB.Preload("Category").Preload("User").Find(&ads).Error; err != nil {
		return nil, err
	}

	return ads, nil
}

func (cr *AdsRepositoryImpl) GetByID(id string) (models.Ads, error) {
	var ads models.Ads

	if err := database.DB.Preload("Category").Preload("User").First(&ads, "id = ?", id).Error; err != nil {
		return models.Ads{}, err
	}

	return ads, nil
}

func (cr *AdsRepositoryImpl) Create(adsInput models.AdsInput) (models.Ads, error) {
	var ads models.Ads = models.Ads{
		Title:       adsInput.Title,
		Description: adsInput.Description,
		StartDate:   adsInput.StartDate,
		EndDate:     adsInput.EndDate,
		CategoryID:  adsInput.CategoryID,
		UserID:      adsInput.UserID,
	}

	result := database.DB.Create(&ads)

	if err := result.Error; err != nil {
		return models.Ads{}, err
	}

	if err := result.Last(&ads).Error; err != nil {
		return models.Ads{}, err
	}

	return ads, nil
}

func (cr *AdsRepositoryImpl) Update(adsInput models.AdsInput, id string) (models.Ads, error) {
	ads, err := cr.GetByID(id)

	if err != nil {
		return models.Ads{}, err
	}

	ads.Title = adsInput.Title
	ads.Description = adsInput.Description
	ads.StartDate = adsInput.StartDate
	ads.EndDate = adsInput.EndDate
	ads.CategoryID = adsInput.CategoryID

	if err := database.DB.Save(&ads).Error; err != nil {
		return models.Ads{}, err
	}

	return ads, nil
}

func (cr *AdsRepositoryImpl) Delete(id string) error {
	ads, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Unscoped().Delete(&ads).Error; err != nil {
		return err
	}

	return nil
}
