package services

import (
	"go-ads-management/middlewares"
	"go-ads-management/models"
	"go-ads-management/repositories"
)

type UserService struct {
	repository repositories.UserRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		repository: repositories.InitUserRepository(),
		jwtAuth:    jwtAuth,
	}
}

func (us *UserService) Register(userInput models.RegisterInput) (models.User, error) {
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput models.LoginInput) (string, error) {
	user, err := us.repository.GetByEmail(userInput)

	if err != nil {
		return "", err
	}

	token, err := us.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
