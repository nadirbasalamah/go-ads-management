package routes

import (
	"go-ads-management/controllers"
	"go-ads-management/middlewares"
	"go-ads-management/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserController(&jwtConfig)

	users := e.Group("/api/v1/users")

	users.POST("/register", userController.Register)
	users.POST("/login", userController.Login)
	users.GET("/info", userController.GetUserInfo,
		echojwt.WithConfig(authMiddlewareConfig),
		middlewares.VerifyToken,
	)
}
