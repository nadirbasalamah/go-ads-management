package repositories

import "go-ads-management/models"

type UserRepository interface {
	Register(userInput models.RegisterInput) (models.User, error)
	GetByEmail(userInput models.LoginInput) (models.User, error)
}
