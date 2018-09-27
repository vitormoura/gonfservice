package main

import "time"

//Message defines the properties of a message
type Message struct {
	From    string
	To      []string
	Subject string
	Message string
	IsHTML  bool
}

//SendMessageResult defines the result of a send message operation
type SendMessageResult struct {
	MessageID string
	Date      time.Time
	Success   bool
	Error     string
}

//AppConfig defines the expected configuration properties of the application
type AppConfig struct {
	SMTP  SMTPConfig
	Debug bool
}

//SMTPConfig defines the expected properties to connect with a SMTP Server
type SMTPConfig struct {
	Host string
	Port int
}
