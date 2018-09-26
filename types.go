package main

import "time"

//Message define the properties of a message
type Message struct {
	From    string
	To      []string
	Subject string
	Message string
	IsHTML  bool
}

//SendMessageResult define the result of a send message operation
type SendMessageResult struct {
	MessageID string
	Date      time.Time
	Success   bool
	Error     string
}
