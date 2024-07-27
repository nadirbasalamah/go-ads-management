package routes

import (
	"go-ads-management/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	//TODO: config log and JWT middleware

	userController := controllers.InitUserController()

	users := e.Group("/api/v1/users")

	users.POST("/register", userController.Register)
	users.POST("/login", userController.Login)
}
