package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	var appConfig *AppConfig
	var err error

	appConfig = new(AppConfig)

	//Standard config
	appConfig.SMTP.Host = "localhost"
	appConfig.SMTP.Port = 25
	appConfig.Debug = false

	//Reading params from command line
	var port = flag.Int("port", 1323, "port of http service")
	var configFilePath = flag.String("config", "app.config.toml", "path to application config file")

	//If a config file was informed, let's try to load it
	if *configFilePath != "" {
		if appConfig, err = loadConfig(*configFilePath); err != nil {
			log.Fatalf("failed to load configuration file: %s", err)
		}
	}

	// Start server
	e, err := createApp(appConfig)

	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
