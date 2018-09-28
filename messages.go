package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/russross/blackfriday"
	"github.com/satori/go.uuid"
	gomail "gopkg.in/mail.v2"
)

//MessageType defines a type to messages
type MessageType int

const (

	//PlainMail plain text e-mail message
	PlainMail MessageType = iota

	//RichFormatMail e-mail with some sort of rich format
	RichFormatMail
)

//Message defines the properties of a message
type Message struct {
	Type    MessageType
	From    string
	To      []string
	CC      []string
	Subject string
	Content string
}

//SendMessageResult defines the result of a send message operation
type SendMessageResult struct {
	MessageID string
	Date      time.Time
	Success   bool
	Errors    []string
}

//MessageSender defines operations to send messages
type MessageSender interface {

	//Send sends the message
	Send(m Message) (SendMessageResult, error)
}

//MailMessageSender sends emails
type MailMessageSender struct {
	Config SMTPConfig
}

//Send sends a message
func (s MailMessageSender) Send(msg Message) (SendMessageResult, error) {

	result := SendMessageResult{
		MessageID: uuid.NewV4().String(),
		Date:      time.Now(),
	}

	//Validation
	result.Errors = validateMailMessage(&msg)

	if len(result.Errors) > 0 {
		return result, nil
	}

	//Preparing content
	m := gomail.NewMessage()

	if msg.Type == PlainMail {

		m.SetBody("text/plain", msg.Content)

	} else if msg.Type == RichFormatMail {

		contentBuffer := bytes.NewBufferString(msg.Content)
		output := blackfriday.Run(contentBuffer.Bytes())

		m.SetBody("text/html", string(output))
	}

	m.SetHeader("From", msg.From)
	m.SetHeader("To", msg.To...)
	m.SetHeader("Subject", msg.Subject)

	if len(msg.CC) > 0 {
		m.SetHeader("Cc", msg.CC...)
	}

	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(s.Config.Host, s.Config.Port, s.Config.Username, s.Config.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return result, fmt.Errorf("error sending message: %s", err)
	}

	result.Success = len(result.Errors) == 0

	return result, nil
}
