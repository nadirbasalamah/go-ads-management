package routes

import (
	"go-ads-management/app/middlewares"
	"go-ads-management/controllers/ads"
	"go-ads-management/controllers/categories"
	"go-ads-management/controllers/users"
	"go-ads-management/drivers/pinata"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	UserController     users.UserController
	CategoryController categories.CategoryController
	AdsController      ads.AdsController
}

func (cl *ControllerList) RegisterRoute(e *echo.Echo) {
	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	e.Use(cl.LoggerMiddleware)

	// TEST ONLY, will be removed later
	e.POST("/upload", func(c echo.Context) error {
		// Retrieve the file from the form input
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		url, err := pinata.UploadFile(file)

		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err,
			})
		}

		return c.JSON(200, echo.Map{
			"message": "upload success",
			"url":     url,
		})
	})

	// user routes
	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", cl.UserController.Register)
	userRoutes.POST("/login", cl.UserController.Login)
	userRoutes.GET(
		"/info",
		cl.UserController.GetUserInfo,
		echojwt.WithConfig(cl.JWTMiddleware),
		middlewares.VerifyToken,
	)

	// category routes
	categoryRoutes := e.Group("/api/v1", echojwt.WithConfig(cl.JWTMiddleware), middlewares.VerifyToken)

	categoryRoutes.GET("/categories", cl.CategoryController.GetAll)
	categoryRoutes.GET("/categories/:id", cl.CategoryController.GetByID)

	categoryRoutes.POST("/categories", cl.CategoryController.Create, middlewares.VerifyAdmin)
	categoryRoutes.PUT("/categories/:id", cl.CategoryController.Update, middlewares.VerifyAdmin)
	categoryRoutes.DELETE("/categories/:id", cl.CategoryController.Delete, middlewares.VerifyAdmin)

	// ads routes
	adsRoutes := e.Group("/api/v1", echojwt.WithConfig(cl.JWTMiddleware), middlewares.VerifyToken)

	adsRoutes.GET("/ads", cl.AdsController.GetAll)
	adsRoutes.GET("/ads/:id", cl.AdsController.GetByID)
	adsRoutes.GET("/ads/category/:category_id", cl.AdsController.GetByCategory)
	adsRoutes.GET("/ads/user", cl.AdsController.GetByUser)
	adsRoutes.GET("/ads/trashed", cl.AdsController.GetTrashed)
	adsRoutes.POST("/ads", cl.AdsController.Create)
	adsRoutes.PUT("/ads/:id", cl.AdsController.Update)
	adsRoutes.DELETE("/ads/:id", cl.AdsController.Delete)
	adsRoutes.POST("/ads/:id", cl.AdsController.Restore)
	adsRoutes.DELETE("/ads/:id/force", cl.AdsController.ForceDelete)
}
