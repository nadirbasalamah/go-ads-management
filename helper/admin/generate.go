package main

import (
	"context"
	"go-ads-management/app/middlewares"
	"go-ads-management/utils"
	"log"

	_driverFactory "go-ads-management/drivers"

	_userUseCase "go-ads-management/businesses/users"

	_dbDriver "go-ads-management/drivers/mysql"
)

type UserInput struct {
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
}

func main() {
	configDB := _dbDriver.DBConfig{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.MigrateDB(db)

	userRepo := _driverFactory.NewUserRepository(db)

	userInput := UserInput{
		CompanyName: utils.GetConfig("ADMIN_COMPANY_NAME"),
		Address:     utils.GetConfig("ADMIN_ADDRESS"),
		Username:    utils.GetConfig("ADMIN_USERNAME"),
		Email:       utils.GetConfig("ADMIN_EMAIL"),
		Password:    utils.GetConfig("ADMIN_PASSWORD"),
	}

	generateAdmin(userRepo, userInput)
}

func generateAdmin(userRepo _userUseCase.Repository, userInput UserInput) {
	customValidator := middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	if err := customValidator.Validate(userInput); err != nil {
		log.Fatalf("invalid input: %v\n", err)
	}

	userReq := &_userUseCase.Domain{
		CompanyName: userInput.CompanyName,
		Address:     userInput.Address,
		Username:    userInput.Username,
		Email:       userInput.Email,
		Password:    userInput.Password,
		Role:        utils.ROLE_ADMIN,
	}

	_, err := userRepo.CreateAdmin(context.TODO(), userReq)

	if err != nil {
		log.Fatalf("failed to create admin account: %v\n", err)
	}

	log.Println("admin account created successfully")
}
