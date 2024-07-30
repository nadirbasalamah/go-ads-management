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

	// user routes
	userController := controllers.InitUserController(&jwtConfig)
	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.GET("/info", userController.GetUserInfo,
		echojwt.WithConfig(authMiddlewareConfig),
		middlewares.VerifyToken,
	)

	// category routes
	categoryController := controllers.InitCategoryController()
	categoryRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	categoryRoutes.GET("/categories", categoryController.GetAll)
	categoryRoutes.GET("/categories/:id", categoryController.GetByID)
	categoryRoutes.POST("/categories", categoryController.Create)
	categoryRoutes.PUT("/categories/:id", categoryController.Update)
	categoryRoutes.DELETE("/categories/:id", categoryController.Delete)

	// ads routes
	adsController := controllers.InitAdsController()
	adsRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	adsRoutes.GET("/ads", adsController.GetAll)
	adsRoutes.GET("/ads/:id", adsController.GetByID)
	adsRoutes.POST("/ads", adsController.Create)
	adsRoutes.PUT("/ads/:id", adsController.Update)
	adsRoutes.DELETE("/ads/:id", adsController.Delete)
}
