package users

import (
	"context"
	"errors"
	"go-ads-management/app/middlewares"
	"strconv"
)

type userUseCase struct {
	userRepository Repository
	jwtConfig      *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtConfig *middlewares.JWTConfig) UseCase {
	return &userUseCase{
		userRepository: repository,
		jwtConfig:      jwtConfig,
	}
}

func (usecase *userUseCase) Register(userReq *Domain) (Domain, error) {
	return usecase.userRepository.Register(userReq)
}

func (usecase *userUseCase) Login(userReq *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(userReq)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtConfig.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usecase *userUseCase) GetUserInfo(ctx context.Context) (Domain, error) {
	claim, err := middlewares.GetUser(ctx)

	if err != nil {
		return Domain{}, errors.New("invalid token")
	}

	userID := strconv.Itoa(claim.ID)

	return usecase.userRepository.GetUserInfo(userID)
}
