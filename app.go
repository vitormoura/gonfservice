package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


//AppContext represents a custom context
type AppContext struct {
	echo.Context

	Config *AppConfig
}

///////////////////////////////////

func createApp(config *AppConfig) (*echo.Echo, error) {

	var (
		app *echo.Echo
		err error
	)

	if config == nil {

		if config, err = loadConfig("app.config.toml"); err != nil {
			return nil, err
		}
	}

	// Echo instance
	app = echo.New()
	app.HideBanner = true

	// Middleware

	if config.Debug {
		app.Use(middleware.Logger())
	}

	app.Use(middleware.Recover())

	app.Use(func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &AppContext{c, config}

			return handler(cc)
		}
	})

	// Routes
	app.POST("/messages", handleNewMessage)

	return app, nil
}