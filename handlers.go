package main

import (
	"fmt"
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
		
	//Message sender
	sender = cc.Sender

	//Reading request payload
	if err = c.Bind(&msg); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("invalid request payload: %s", err))
	}

	//Sending message
	if result, err = sender.Send(msg); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("send message failed due to an internal error: %s", err))
	}
	
	if len(result.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, result)
	}

	return c.JSON(http.StatusOK, result)
}