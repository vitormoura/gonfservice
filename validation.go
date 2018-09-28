package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func validateMailMessage(msg *Message) (errors []string) {

	//Validate message
	if errMsg := validateMailAddressFormat(msg.From); errMsg != "" {
		errors = append(errors, errMsg)
	}

	//To
	if len(msg.To) == 0 {
		errors = append(errors, fmt.Sprintf("message without destination"))
	} else {
		validateMailAddressList(msg.To, &errors)
	}

	//CC
	if len(msg.CC) > 0 {
		validateMailAddressList(msg.CC, &errors)
	}

	//Content
	if msg.Content == "" {
		errors = append(errors, fmt.Sprintf("message without content"))
	}

	//Supported types
	if msg.Type != PlainMail && msg.Type != RichFormatMail {
		errors = append(errors, fmt.Sprintf("invalid message type: '%d'", msg.Type))
	}

	return
}

func validateMailAddressList(emails []string, errors *[]string) {

	for _, address := range emails {
		if errMsg := validateMailAddressFormat(address); errMsg != "" {
			*errors = append(*errors, errMsg)
		}
	}

	return
}

func validateMailAddressFormat(address string) string {

	if err := checkmail.ValidateFormat(address); err != nil {
		return fmt.Sprintf("invalid email address: '%s'", address)
	}

	return ""
}
