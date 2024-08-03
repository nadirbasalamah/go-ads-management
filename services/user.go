package services

import (
	"go-ads-management/models"
	"go-ads-management/repositories"
	"go-ads-management/utils"
)

type UserService struct {
	Repository repositories.UserRepository
	jwtOptions models.JWTOptions
}

func InitUserService(jwtOptions models.JWTOptions) UserService {
	return UserService{
		Repository: repositories.InitUserRepository(),
		jwtOptions: jwtOptions,
	}
}

func (us *UserService) Register(userInput models.RegisterInput) (models.User, error) {
	return us.Repository.Register(userInput)
}

func (us *UserService) Login(userInput models.LoginInput) (string, error) {
	user, err := us.Repository.GetByEmail(userInput)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(user.ID), us.jwtOptions)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUserInfo(id string) (models.User, error) {
	return us.Repository.GetUserInfo(id)
}
