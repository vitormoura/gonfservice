package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleNewMessage(c echo.Context) error {

	var (
		result SendMessageResult
		sender MessageSender
		err error
		cc = c.(*AppContext)
		msg Message
	)
	
	mailSender := MailMessageSender {
		Config: cc.Config.SMTP,
	}

	sender = &mailSender	

	//Reading request payload
	if err = c.Bind(&msg); err != nil {
		return c.String(http.StatusBadRequest, "invalid request payload")
	}

	//Sending message
	if result, err = sender.Send(msg); err != nil {
		return c.String(http.StatusInternalServerError, "send message failed due to an internal error")
	}
	
	if len(result.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, result)
	}

	return c.JSON(http.StatusOK, result)
}