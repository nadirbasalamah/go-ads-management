package ads

import (
	"go-ads-management/businesses/ads"

	"gorm.io/gorm"
)

type adsRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) ads.Repository {
	return &adsRepository{
		conn: conn,
	}
}

func (ar *adsRepository) GetAll() (*gorm.DB, error) {
	stmt := ar.conn.Joins("Category").Joins("User").Model(Ads{})

	return stmt, nil
}

func (ar *adsRepository) GetByID(id string) (ads.Domain, error) {
	var adsData Ads

	if err := ar.conn.Preload("Category").Preload("User").First(&adsData, "id = ?", id).Error; err != nil {
		return ads.Domain{}, err
	}

	return adsData.ToDomain(), nil
}

func (ar *adsRepository) Create(adsReq *ads.Domain) (ads.Domain, error) {
	record := FromDomain(adsReq)

	result := ar.conn.Create(&record)

	if err := result.Error; err != nil {
		return ads.Domain{}, err
	}

	if err := result.Preload("Category").Preload("User").Last(&record).Error; err != nil {
		return ads.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ar *adsRepository) Update(adsReq *ads.Domain, id string) (ads.Domain, error) {
	adsData, err := ar.GetByID(id)

	if err != nil {
		return ads.Domain{}, err
	}

	updatedAds := FromDomain(&adsData)

	updatedAds.Title = adsReq.Title
	updatedAds.Description = adsReq.Description
	updatedAds.StartDate = adsReq.StartDate
	updatedAds.EndDate = adsReq.EndDate
	updatedAds.CategoryID = adsReq.CategoryID

	if err := ar.conn.Save(&updatedAds).Error; err != nil {
		return ads.Domain{}, err
	}

	return updatedAds.ToDomain(), nil
}

func (ar *adsRepository) Delete(id string) error {
	adsData, err := ar.GetByID(id)

	if err != nil {
		return err
	}

	deletedAds := FromDomain(&adsData)

	if err := ar.conn.Delete(&deletedAds).Error; err != nil {
		return err
	}

	return nil
}

func (ar *adsRepository) Restore(id string) (ads.Domain, error) {
	var adsData Ads

	err := ar.conn.Unscoped().First(&adsData, "id = ?", id).Error

	if err != nil {
		return ads.Domain{}, err
	}

	adsData.DeletedAt = gorm.DeletedAt{}

	err = ar.conn.Unscoped().Save(&adsData).Error

	if err != nil {
		return ads.Domain{}, err
	}

	return adsData.ToDomain(), nil
}

func (ar *adsRepository) ForceDelete(id string) error {
	adsData, err := ar.GetByID(id)

	if err != nil {
		return err
	}

	deletedAds := FromDomain(&adsData)

	if err := ar.conn.Unscoped().Delete(&deletedAds).Error; err != nil {
		return err
	}

	return nil
}
