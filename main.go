package main

import (
	"go-notes/config"
	"go-notes/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.Migrate()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	e.Start(":8080")
}
