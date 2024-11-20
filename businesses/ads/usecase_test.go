package ads_test

import (
	"context"
	"errors"
	"go-ads-management/businesses/ads"
	_adsMock "go-ads-management/businesses/ads/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	adsRepository _adsMock.Repository
	adsUseCase    ads.UseCase
	ctx           context.Context
	queryResult   *gorm.DB
)

func TestMain(m *testing.M) {
	adsUseCase = ads.NewAdsUseCase(&adsRepository)
	ctx = context.TODO()
	queryResult = &gorm.DB{}

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		adsRepository.On("GetAll", ctx).Return(queryResult, nil).Once()

		result, err := adsUseCase.GetAll(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		adsRepository.On("GetAll", ctx).Return(queryResult, errors.New("error")).Once()

		result, err := adsUseCase.GetAll(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		adsRepository.On("GetByID", ctx, 1).Return(ads.Domain{}, nil).Once()

		result, err := adsUseCase.GetByID(ctx, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		adsRepository.On("GetByID", ctx, 0).Return(ads.Domain{}, errors.New("whoops")).Once()

		result, err := adsUseCase.GetByID(ctx, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByCategory(t *testing.T) {
	t.Run("GetByCategory | Valid", func(t *testing.T) {
		adsRepository.On("GetByCategory", ctx, 1).Return(queryResult, nil).Once()

		result, err := adsUseCase.GetByCategory(ctx, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByCategory | Invalid", func(t *testing.T) {
		adsRepository.On("GetByCategory", ctx, 0).Return(queryResult, errors.New("whoops")).Once()

		result, err := adsUseCase.GetByCategory(ctx, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByUser(t *testing.T) {
	t.Run("GetByUser | Valid", func(t *testing.T) {
		adsRepository.On("GetByUser", ctx).Return(queryResult, nil).Once()

		result, err := adsUseCase.GetByUser(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByUser | Invalid", func(t *testing.T) {
		adsRepository.On("GetByUser", ctx).Return(queryResult, errors.New("whoops")).Once()

		result, err := adsUseCase.GetByUser(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetTrashed(t *testing.T) {
	t.Run("GetTrashed | Valid", func(t *testing.T) {
		adsRepository.On("GetTrashed", ctx).Return(queryResult, nil).Once()

		result, err := adsUseCase.GetTrashed(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetTrashed | Invalid", func(t *testing.T) {
		adsRepository.On("GetTrashed", ctx).Return(queryResult, errors.New("whoops")).Once()

		result, err := adsUseCase.GetTrashed(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		adsRepository.On("Create", ctx, &ads.Domain{}).Return(ads.Domain{}, nil).Once()

		result, err := adsUseCase.Create(ctx, &ads.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		adsRepository.On("Create", ctx, &ads.Domain{}).Return(ads.Domain{}, errors.New("whoops")).Once()

		result, err := adsUseCase.Create(ctx, &ads.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		adsRepository.On("Update", ctx, &ads.Domain{}, 1).Return(ads.Domain{}, nil).Once()

		result, err := adsUseCase.Update(ctx, &ads.Domain{}, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		adsRepository.On("Update", ctx, &ads.Domain{}, 0).Return(ads.Domain{}, errors.New("whoops")).Once()

		result, err := adsUseCase.Update(ctx, &ads.Domain{}, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		adsRepository.On("Delete", ctx, 1).Return(nil).Once()

		err := adsUseCase.Delete(ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		adsRepository.On("Delete", ctx, 0).Return(errors.New("whoops")).Once()

		err := adsUseCase.Delete(ctx, 0)

		assert.NotNil(t, err)
	})
}

func TestRestore(t *testing.T) {
	t.Run("Restore | Valid", func(t *testing.T) {
		adsRepository.On("Restore", ctx, 1).Return(ads.Domain{}, nil).Once()

		result, err := adsUseCase.Restore(ctx, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Restore | Invalid", func(t *testing.T) {
		adsRepository.On("Restore", ctx, 0).Return(ads.Domain{}, errors.New("whoops")).Once()

		result, err := adsUseCase.Restore(ctx, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestForceDelete(t *testing.T) {
	t.Run("ForceDelete | Valid", func(t *testing.T) {
		adsRepository.On("ForceDelete", ctx, 1).Return(nil).Once()

		err := adsUseCase.ForceDelete(ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("ForceDelete | Invalid", func(t *testing.T) {
		adsRepository.On("ForceDelete", ctx, 0).Return(errors.New("whoops")).Once()

		err := adsUseCase.ForceDelete(ctx, 0)

		assert.NotNil(t, err)
	})
}
