package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

func handleNewMessage(c echo.Context) error {

	msg := new(Message)

	if err := c.Bind(msg); err != nil {
		return err
	}
	
	resultStatusCode := http.StatusCreated

	result := SendMessageResult{
		MessageID: uuid.NewV4().String(),
		Success:   true,
		Date:      time.Now(),
	}

	//Validate message
	if msg.From == "" {
		result.Success = false
		result.Error = MsgInvalidMessage
		resultStatusCode = http.StatusBadRequest
	}

	return c.JSON(resultStatusCode, result)
}
