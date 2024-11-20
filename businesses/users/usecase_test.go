package users_test

import (
	"context"
	"errors"
	"go-ads-management/app/middlewares"
	"go-ads-management/businesses/users"
	_userMock "go-ads-management/businesses/users/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.Repository
	userUseCase    users.UseCase
	ctx            context.Context
)

func TestMain(m *testing.M) {
	userUseCase = users.NewUserUseCase(&userRepository, &middlewares.JWTConfig{})
	ctx = context.TODO()

	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", ctx, &users.Domain{}).Return(users.Domain{}, nil).Once()

		result, err := userUseCase.Register(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.Register(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &users.Domain{}).Return(users.Domain{}, nil).Once()

		result, err := userUseCase.Login(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Login | Invalid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.Login(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", ctx).Return(users.Domain{}, nil).Once()

		result, err := userUseCase.GetUserInfo(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", ctx).Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.GetUserInfo(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
