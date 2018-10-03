package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


//AppContext represents a custom context
type AppContext struct {
	echo.Context
	
	Sender MessageSender
}

///////////////////////////////////

func createApp(senderCreator MessageSenderCreator , debug bool) (*echo.Echo, error) {

	var (
		app *echo.Echo
	)
	 
	// Echo instance
	app = echo.New()
	app.HideBanner = true

	// Middleware

	if debug {
		app.Use(middleware.Logger())
	}

	app.Use(middleware.Recover())

	// Middleware invoked every request
	app.Use(func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &AppContext{
				c, 
				senderCreator(), 
			}
						
			return handler(cc)
		}
	})

	// Routes
	app.POST("/messages", handleNewMessage)

	return app, nil
}