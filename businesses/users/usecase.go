package users

import (
	"context"
	"go-ads-management/app/middlewares"
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

func (usecase *userUseCase) Register(ctx context.Context, userReq *Domain) (Domain, error) {
	return usecase.userRepository.Register(ctx, userReq)
}

func (usecase *userUseCase) Login(ctx context.Context, userReq *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userReq)

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
	return usecase.userRepository.GetUserInfo(ctx)
}
