package main

import (
	"go-ads-management/database"
	"go-ads-management/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.Migrate()

	//TODO: start the server
	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
