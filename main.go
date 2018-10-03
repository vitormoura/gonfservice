package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	var (
		appConfig *AppConfig
		err       error
	)

	appConfig = new(AppConfig)

	//Standard config
	appConfig.SMTP.Host = "localhost"
	appConfig.SMTP.Port = 25
	appConfig.SMTP.Timeout = 60
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

	//Mail sender creator
	defaultSenderCreator := func() MessageSender {
		return MailMessageSender{appConfig.SMTP}
	}

	// Start server
	e, err := createApp(defaultSenderCreator, appConfig.Debug)

	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
