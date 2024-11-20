package categories_test

import (
	"context"
	"errors"
	"go-ads-management/businesses/categories"
	_categoryMock "go-ads-management/businesses/categories/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository _categoryMock.Repository
	categoryUseCase    categories.UseCase
	ctx                context.Context
)

func TestMain(m *testing.M) {
	categoryUseCase = categories.NewCategoryUseCase(&categoryRepository)
	ctx = context.TODO()

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]categories.Domain{}, nil).Once()

		result, err := categoryUseCase.GetAll(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]categories.Domain{}, errors.New("error")).Once()

		result, err := categoryUseCase.GetAll(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, 1).Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.GetByID(ctx, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, 0).Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.GetByID(ctx, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, &categories.Domain{}).Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.Create(ctx, &categories.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, &categories.Domain{}).Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.Create(ctx, &categories.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, &categories.Domain{}, 1).Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.Update(ctx, &categories.Domain{}, 1)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, &categories.Domain{}, 0).Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.Update(ctx, &categories.Domain{}, 0)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, 1).Return(nil).Once()

		err := categoryUseCase.Delete(ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, 0).Return(errors.New("whoops")).Once()

		err := categoryUseCase.Delete(ctx, 0)

		assert.NotNil(t, err)
	})
}
