package main

import (
	"go-ads-management/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	//TODO: connect to the database
	database.InitDB()

	//TODO: start the server
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
