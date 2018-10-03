package main

import (
	"bytes"
	"crypto/tls"
"fmt"
"time"
	"github.com/russross/blackfriday"
	"github.com/satori/go.uuid"
	"github.com/go-mail/mail"
)

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
	m := mail.NewMessage()

	if msg.Type == PlainMail {

		m.SetBody("text/plain", msg.Content)

	} else if msg.Type == RichFormatMail {

		contentBuffer := bytes.NewBufferString(msg.Content)
		output := blackfriday.Run(contentBuffer.Bytes())

		m.SetBody("text/html", string(output))
	}
	
	m.SetHeader("Subject", msg.Subject)

	if len(msg.CC) > 0 {
		m.SetHeader("Cc", msg.CC...)
	}

	//m.Attach("/home/Alex/lolcat.jpg")

	var (
		sender mail.SendCloser
		err    error
	)

	d := mail.NewDialer(s.Config.Host, s.Config.Port, s.Config.Username, s.Config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	d.Timeout = time.Second * 30

	//
	if sender, err = d.Dial(); err != nil {
		return result, fmt.Errorf("error dialing remote smtp server: %s", err)
	}

	defer sender.Close()

	// Send the email
	if err := sender.Send(msg.From, msg.To, m); err != nil {
		return result, fmt.Errorf("error sending message: %s", err)
	}

	result.Success = len(result.Errors) == 0

	return result, nil
}
