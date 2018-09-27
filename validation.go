package main

import (
	"fmt"
)

const (

	//MsgInvalidMessage informs an invalid message
	MsgInvalidMessage = "Invalid message"

	//MsgErrorSendingMessage informs an error sending messages
	MsgErrorSendingMessage = "Error sending message"
)

///////////////////////////////////

func validateMailMessage(msg *Message) (errors []string) {

	//Validate message
	if msg.From == "" {
		errors = append(errors, MsgInvalidMessage)
	}

	//To field empty
	if len(msg.To) == 0 {
		errors = append(errors, fmt.Sprintf("message without "))
	}

	//Content
	if msg.Content == "" {
		errors = append(errors, fmt.Sprintf("message without content"))
	}

	return
}
