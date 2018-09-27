package main

import (
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	toml "github.com/pelletier/go-toml"
)

func main() {

	// Start server
	e, err := createApp(nil)

	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

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
	//e.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.POST("/messages", handleNewMessage)
	
	return app, nil
}

func loadConfig(path string) (*AppConfig, error) {
	
	var (

		err error
		tomlData *toml.Tree
		config AppConfig
	)

	// Load application config file
	if tomlData, err = toml.LoadFile(path); err != nil {
		return nil, err
	}

	if err = tomlData.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
