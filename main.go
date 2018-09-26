package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/satori/go.uuid"
	toml "github.com/pelletier/go-toml"
)

func main() {
	
	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func (c echo.Context) error {
		u1 := uuid.NewV4()
		return c.String(http.StatusOK, u1.String())
	})

	// Load application config file
	toml.LoadFile("app.config.toml")

	// Start server

	e.Logger.Fatal(e.Start(":1323"))
}