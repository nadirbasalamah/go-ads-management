package repositories_test

import (
	"errors"
	"go-ads-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetAllAds(t *testing.T) {
	t.Run("Get All Ads | Valid", func(t *testing.T) {
		adsRepository.On("GetAll").Return(&gorm.DB{}, nil).Once()

		result, err := adsService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Ads | Invalid", func(t *testing.T) {
		adsRepository.On("GetAll").Return(&gorm.DB{}, errors.New("error")).Once()

		result, err := adsService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetAdsByID(t *testing.T) {
	t.Run("Get Ads By ID | Valid", func(t *testing.T) {
		adsRepository.On("GetByID", "1").Return(models.Ads{}, nil).Once()

		result, err := adsService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Ads By ID | Invalid", func(t *testing.T) {
		adsRepository.On("GetByID", "-1").Return(models.Ads{}, errors.New("error")).Once()

		result, err := adsService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateAds(t *testing.T) {
	t.Run("Create Ads | Valid", func(t *testing.T) {
		adsRepository.On("Create", models.AdsInput{}).Return(models.Ads{}, nil).Once()

		result, err := adsService.Create(models.AdsInput{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Ads | Invalid", func(t *testing.T) {
		adsRepository.On("Create", models.AdsInput{}).Return(models.Ads{}, errors.New("error")).Once()

		result, err := adsService.Create(models.AdsInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateAds(t *testing.T) {
	t.Run("Update Ads | Valid", func(t *testing.T) {
		adsRepository.On("Update", models.AdsInput{}, "1").Return(models.Ads{}, nil).Once()

		result, err := adsService.Update(models.AdsInput{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Ads | Invalid", func(t *testing.T) {
		adsRepository.On("Update", models.AdsInput{}, "-1").Return(models.Ads{}, errors.New("error")).Once()

		result, err := adsService.Update(models.AdsInput{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteAds(t *testing.T) {
	t.Run("Delete Ads | Valid", func(t *testing.T) {
		adsRepository.On("Delete", "1").Return(nil).Once()

		err := adsService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Ads | Invalid", func(t *testing.T) {
		adsRepository.On("Delete", "-1").Return(errors.New("error")).Once()

		err := adsService.Delete("-1")

		assert.NotNil(t, err)
	})
}

func TestRestoreAds(t *testing.T) {
	t.Run("Restore Ads | Valid", func(t *testing.T) {
		adsRepository.On("Restore", "1").Return(models.Ads{}, nil).Once()

		result, err := adsService.Restore("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Restore Ads | Invalid", func(t *testing.T) {
		adsRepository.On("Restore", "-1").Return(models.Ads{}, errors.New("error")).Once()

		result, err := adsService.Restore("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestForceDeleteAds(t *testing.T) {
	t.Run("Force Delete Ads | Valid", func(t *testing.T) {
		adsRepository.On("ForceDelete", "1").Return(nil).Once()

		err := adsService.ForceDelete("1")

		assert.Nil(t, err)
	})

	t.Run("Force Delete Ads | Invalid", func(t *testing.T) {
		adsRepository.On("ForceDelete", "-1").Return(errors.New("error")).Once()

		err := adsService.ForceDelete("-1")

		assert.NotNil(t, err)
	})
}
