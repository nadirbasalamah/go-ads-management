package main

import (
	"context"
	"go-ads-management/app/middlewares"
	"go-ads-management/utils"
	"log"
	"os"
	"os/exec"
	"runtime"

	_driverFactory "go-ads-management/drivers"

	_userUseCase "go-ads-management/businesses/users"

	_dbDriver "go-ads-management/drivers/mysql"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
}

func main() {
	userInput := UserInput{
		CompanyName: utils.GetConfig("ADMIN_COMPANY_NAME"),
		Address:     utils.GetConfig("ADMIN_ADDRESS"),
		Username:    utils.GetConfig("ADMIN_USERNAME"),
		Email:       utils.GetConfig("ADMIN_EMAIL"),
		Password:    utils.GetConfig("ADMIN_PASSWORD"),
	}

	if utils.GetConfig("APP_MODE") != "production" {
		createAdmin(userInput)
	} else {
		execInitAdmin(userInput)
	}
}

func createAdmin(userInput UserInput) {
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

	validateRequest(userInput)

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

func execInitAdmin(userInput UserInput) {
	validateRequest(userInput)

	bs, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error occurred when creating password: %v\n", err)
	}

	var password string = string(bs)

	databaseName := utils.GetConfig("DB_NAME")
	databaseUsername := utils.GetConfig("DB_USERNAME")
	databasePassword := utils.GetConfig("DB_PASSWORD")
	companyName := userInput.CompanyName
	address := userInput.Address
	username := userInput.Username
	email := userInput.Email

	// Determine the appropriate command to execute the script based on OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// On Windows, use "bash" to execute the shell script
		cmd = exec.Command(
			"bash",
			"./scripts/init_admin.sh",
			databaseName,
			databaseUsername,
			databasePassword,
			companyName,
			address,
			username,
			email,
			password,
		)
	} else {
		// On Unix-based systems, execute the init_admin directly
		cmd = exec.Command(
			"./scripts/init_admin.sh",
			databaseName,
			databaseUsername,
			databasePassword,
			companyName,
			address,
			username,
			email,
			password,
		)
	}

	// Set the environment variables if needed
	cmd.Env = os.Environ()

	// Capture output and errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Script execution failed: %v\nOutput: %s", err, output)
	}

	// Print the output from the script
	log.Printf("Script executed successfully:\n%s", output)

	log.Println("admin created successfully")
}

func validateRequest(userInput UserInput) {
	customValidator := middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	if err := customValidator.Validate(userInput); err != nil {
		log.Fatalf("invalid input: %v\n", err)
	}
}
