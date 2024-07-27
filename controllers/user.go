package controllers

import (
	"go-ads-management/models"
	"go-ads-management/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController() UserController {
	return UserController{
		service: services.InitUserService(),
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var userInput models.RegisterInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	user, err := uc.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:  "success",
		Message: "user registered",
		Data:    user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput models.LoginInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := uc.service.Login(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}
