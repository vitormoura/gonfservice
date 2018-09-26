package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	toml "github.com/pelletier/go-toml"
)

func main() {

	// Start server
	e := createApp()
	e.Logger.Fatal(e.Start(":1323"))
}

func createApp() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/messages", handleNewMessage)

	// Load application config file
	toml.LoadFile("app.config.toml")

	return e
}