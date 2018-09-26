package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	toml "github.com/pelletier/go-toml"
	"github.com/satori/go.uuid"
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

func handleNewMessage(c echo.Context) error {

	uniqueID := uuid.NewV4()

	result := SendMessageResult{
		MessageID: uniqueID.String(),
		Success:   true,
		Date:      time.Now(),
	}

	return c.JSON(http.StatusCreated, result) 
}
