package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	gomail "gopkg.in/mail.v2"
)

func handleNewMessage(c echo.Context) error {

	cc := c.(*AppContext)
	msg := new(Message)

	if err := c.Bind(msg); err != nil {
		return err
	}
	
	result := SendMessageResult{
		MessageID: uuid.NewV4().String(),
		Success:   true,
		Date:      time.Now(),
	}

	result.Errors = validateMailMessage(msg)

	if len(result.Errors) > 0 {
		
		result.Success = false

		return cc.JSON(http.StatusBadRequest, result)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", msg.From)
	m.SetHeader("To", msg.To...)
	m.SetHeader("Subject", msg.Subject)

	if len(msg.CC) > 0 {
		m.SetHeader("Cc", msg.CC...)
	}

	m.SetBody("text/plain", msg.Content)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(cc.Config.SMTP.Host, cc.Config.SMTP.Port, cc.Config.SMTP.Username, cc.Config.SMTP.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		result.Errors = append(result.Errors, MsgErrorSendingMessage)
	}
	
	return c.JSON(http.StatusOK, result)
}