package main

import toml "github.com/pelletier/go-toml"

//AppConfig defines the expected configuration properties of the application
type AppConfig struct {
	SMTP  SMTPConfig
	Debug bool
}

//SMTPConfig defines the expected properties to connect with a SMTP Server
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Timeout  int
}

///////////////////////////////////

func loadConfig(path string) (*AppConfig, error) {

	var (
		err      error
		tomlData *toml.Tree
		config   AppConfig
	)

	// Load application config file
	if tomlData, err = toml.LoadFile(path); err != nil {
		return nil, err
	}

	// Parse content into config object
	if err = tomlData.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
