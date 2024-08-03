package repositories_test

import (
	"errors"
	"go-ads-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories(t *testing.T) {
	t.Run("Get All Categories | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{}, nil).Once()

		result, err := categoryService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Categories | Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{}, errors.New("error")).Once()

		result, err := categoryService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("Get Category By ID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", "1").Return(models.Category{}, nil).Once()

		result, err := categoryService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Category By ID | Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", "-1").Return(models.Category{}, errors.New("error")).Once()

		result, err := categoryService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("Create Category | Valid", func(t *testing.T) {
		categoryRepository.On("Create", models.CategoryInput{}).Return(models.Category{}, nil).Once()

		result, err := categoryService.Create(models.CategoryInput{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Category | Invalid", func(t *testing.T) {
		categoryRepository.On("Create", models.CategoryInput{}).Return(models.Category{}, errors.New("error")).Once()

		result, err := categoryService.Create(models.CategoryInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Update Category | Valid", func(t *testing.T) {
		categoryRepository.On("Update", models.CategoryInput{}, "1").Return(models.Category{}, nil).Once()

		result, err := categoryService.Update(models.CategoryInput{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Category | Invalid", func(t *testing.T) {
		categoryRepository.On("Update", models.CategoryInput{}, "-1").Return(models.Category{}, errors.New("error")).Once()

		result, err := categoryService.Update(models.CategoryInput{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("Delete Category | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", "1").Return(nil).Once()

		err := categoryService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Category | Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", "-1").Return(errors.New("error")).Once()

		err := categoryService.Delete("-1")

		assert.NotNil(t, err)
	})
}
