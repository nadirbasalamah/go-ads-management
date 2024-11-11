package ads

import (
	"context"
	"errors"
	"go-ads-management/app/middlewares"
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

func (ar *adsRepository) GetAll(ctx context.Context) (*gorm.DB, error) {
	stmt := ar.conn.WithContext(ctx).Joins("Category").Joins("User").Model(Ads{})

	return stmt, nil
}

func (ar *adsRepository) GetByID(ctx context.Context, id string) (ads.Domain, error) {
	var adsData Ads

	if err := ar.conn.WithContext(ctx).Preload("Category").Preload("User").First(&adsData, "id = ?", id).Error; err != nil {
		return ads.Domain{}, err
	}

	return adsData.ToDomain(), nil
}

func (ar *adsRepository) GetByCategory(ctx context.Context, categoryID string) (*gorm.DB, error) {
	stmt := ar.conn.WithContext(ctx).Joins("Category").Joins("User").Where("category_id = ?", categoryID).Model(Ads{})

	return stmt, nil
}

func (ar *adsRepository) GetByUser(ctx context.Context) (*gorm.DB, error) {
	userID, err := middlewares.GetUserID(ctx)

	if err != nil {
		return nil, errors.New("invalid token")
	}

	stmt := ar.conn.WithContext(ctx).Joins("Category").Joins("User").Where("user_id = ?", userID).Model(Ads{})

	return stmt, nil
}

func (ar *adsRepository) GetTrashed(ctx context.Context) (*gorm.DB, error) {
	stmt := ar.conn.WithContext(ctx).Unscoped().Joins("Category").Joins("User").Where("ads.deleted_at IS NOT NULL").Model(Ads{})

	return stmt, nil
}

func (ar *adsRepository) Create(ctx context.Context, adsReq *ads.Domain) (ads.Domain, error) {
	userID, err := middlewares.GetUserID(ctx)

	if err != nil {
		return ads.Domain{}, errors.New("invalid token")
	}

	adsReq.UserID = uint(userID)

	record := FromDomain(adsReq)

	result := ar.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return ads.Domain{}, err
	}

	if err := result.WithContext(ctx).Preload("Category").Preload("User").Last(&record).Error; err != nil {
		return ads.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ar *adsRepository) Update(ctx context.Context, adsReq *ads.Domain, id string) (ads.Domain, error) {
	adsData, err := ar.GetByID(ctx, id)

	if err != nil {
		return ads.Domain{}, err
	}

	if !verifyAdsOwner(ctx, adsData) {
		return ads.Domain{}, errors.New("operation not permitted")
	}

	updatedAds := FromDomain(&adsData)

	updatedAds.Title = adsReq.Title
	updatedAds.Description = adsReq.Description
	updatedAds.StartDate = adsReq.StartDate
	updatedAds.EndDate = adsReq.EndDate
	updatedAds.CategoryID = adsReq.CategoryID
	updatedAds.MediaURL = adsReq.MediaURL
	updatedAds.MediaCID = adsReq.MediaCID
	updatedAds.MediaID = adsReq.MediaID

	if err := ar.conn.WithContext(ctx).Save(&updatedAds).Error; err != nil {
		return ads.Domain{}, err
	}

	return updatedAds.ToDomain(), nil
}

func (ar *adsRepository) Delete(ctx context.Context, id string) error {
	adsData, err := ar.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if !verifyAdsOwner(ctx, adsData) {
		return errors.New("operation not permitted")
	}

	deletedAds := FromDomain(&adsData)

	if err := ar.conn.WithContext(ctx).Delete(&deletedAds).Error; err != nil {
		return err
	}

	return nil
}

func (ar *adsRepository) Restore(ctx context.Context, id string) (ads.Domain, error) {
	var adsData Ads

	err := ar.conn.WithContext(ctx).Unscoped().Preload("User").First(&adsData, "id = ?", id).Error

	if err != nil {
		return ads.Domain{}, err
	}

	if !verifyAdsOwner(ctx, adsData.ToDomain()) {
		return ads.Domain{}, errors.New("operation not permitted")
	}

	adsData.DeletedAt = gorm.DeletedAt{}

	err = ar.conn.WithContext(ctx).Unscoped().Save(&adsData).Error

	if err != nil {
		return ads.Domain{}, err
	}

	return adsData.ToDomain(), nil
}

func (ar *adsRepository) ForceDelete(ctx context.Context, id string) error {
	adsData, err := ar.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if !verifyAdsOwner(ctx, adsData) {
		return errors.New("operation not permitted")
	}

	deletedAds := FromDomain(&adsData)

	if err := ar.conn.WithContext(ctx).Unscoped().Delete(&deletedAds).Error; err != nil {
		return err
	}

	return nil
}

func verifyAdsOwner(ctx context.Context, adsData ads.Domain) bool {
	userID, err := middlewares.GetUserID(ctx)

	if err != nil {
		return false
	}

	return adsData.UserID == uint(userID)
}
