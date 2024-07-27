package services

import (
	"go-ads-management/models"
	"go-ads-management/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func InitUserService() UserService {
	return UserService{
		repository: repositories.InitUserRepository(),
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

	//TODO: generate JWT token

	return user.Email, nil
}
