package repositories_test

import (
	"errors"
	"go-ads-management/models"
	"go-ads-management/repositories/mocks"
	"go-ads-management/services"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository mocks.UserRepository
	userService    services.UserService

	categoryRepository mocks.CategoryRepository
	categoryService    services.CategoryService

	adsRepository mocks.AdsRepository
	adsService    services.AdsService
)

func TestMain(m *testing.M) {
	userService = services.InitUserService(models.JWTOptions{})
	userService.Repository = &userRepository

	categoryService = services.InitCategoryService()
	categoryService.Repository = &categoryRepository

	adsService = services.InitAdsService()
	adsService.Repository = &adsRepository

	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterInput{}).Return(models.User{}, nil).Once()

		result, err := userService.Register(models.RegisterInput{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterInput{}).Return(models.User{}, errors.New("error")).Once()

		result, err := userService.Register(models.RegisterInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByEmail(t *testing.T) {
	t.Run("GetByEmail | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", models.LoginInput{}).Return(models.User{}, nil).Once()

		result, err := userService.Login(models.LoginInput{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByEmail | Invalid", func(t *testing.T) {
		userRepository.On("GetByEmail", models.LoginInput{}).Return(models.User{}, errors.New("error")).Once()

		result, err := userService.Login(models.LoginInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "1").Return(models.User{}, nil).Once()

		result, err := userService.GetUserInfo("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "-1").Return(models.User{}, errors.New("error")).Once()

		result, err := userService.GetUserInfo("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
