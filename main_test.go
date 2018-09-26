package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	app    *echo.Echo
	server *httptest.Server
	client *http.Client
)

func init() {
	app = createApp()
	server = httptest.NewServer(app)
	client = server.Client()
}

//////////////////////////////

func Test_Sending_ValidRequest_ResultOK(t *testing.T) {

	result, ok := testPostNewMessage(t, Message{
		From:    "beltrano@mail.com",
		Message: "Um exemplo",
		To: []string{
			"fulano@mail.com",
		},
		IsHTML: false,
	}, http.StatusCreated)

	if ok {
		assertValidSendMessageResult(t, result)
	}
}

func Test_Sending_RequestWithoutFromField_ResultBadRequest(t *testing.T) {

	result, ok := testPostNewMessage(t, Message{
		From:    "",
		Message: "Um exemplo",
		To: []string{
			"fulano@mail.com",
		},
		IsHTML: false,
	}, http.StatusBadRequest)

	if ok {
		assertErrorSendMessageResult(t, result, MsgInvalidMessage)
	}
}

//////////////////////////////

func assertErrorSendMessageResult(t *testing.T, result SendMessageResult, expectedMessage string) {
	assert.False(t, result.Success)
	assert.Equal(t, expectedMessage, result.Error)
}

func assertValidSendMessageResult(t *testing.T, result SendMessageResult) {
	assert.True(t, result.Success)
	assert.True(t, result.MessageID != "")
	assert.True(t, result.Error == "")
	assert.False(t, result.Date.IsZero())
}

func testPostNewMessage(t *testing.T, msg Message, expectedStatusCode int) (result SendMessageResult, ok bool) {

	ok = false
	result = SendMessageResult{}
	msgBytes, _ := json.Marshal(msg)
	buffer := bytes.NewBuffer(msgBytes)

	resp, err := client.Post(server.URL+"/messages", "application/json", buffer)

	if assert.Nil(t, err) {

		if assert.Equal(t, expectedStatusCode, resp.StatusCode) {
			defer resp.Body.Close()

			respBuffer := bytes.Buffer{}
			n, _ := respBuffer.ReadFrom(resp.Body)

			if assert.True(t, n > 0) {
				err := json.Unmarshal(respBuffer.Bytes(), &result)

				assert.Nil(t, err)

				ok = true
			}
		}
	}

	return
}
