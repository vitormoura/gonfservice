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
	config AppConfig
)

func init() {

	config.SMTP.Host = "localhost"
	config.SMTP.Port = 25
	config.Debug = false
	
	app, _ = createApp(createDefaultMailSender, config.Debug)
	server = httptest.NewServer(app)
	client = server.Client()
}

///////////////////////////////////

func Test_Sending_ValidRequest_ResultOK(t *testing.T) {

	result, ok := testPostNewMessage(t, Message{
		From: "beltrano@mail.com",
		Content: `# An important message
		- Item 1
		- Item 2
		- Item 3
		`,
		Subject: "Testing sending valid request",
		CC: []string{
			"sicrano@mail.com",
			"sicrano_2@mail.com",
			"sicrano_3@mail.com",
		},
		To: []string{
			"fulano@mail.com",
		},
		Type: RichFormatMail,
	}, http.StatusOK)

	assert.True(t, ok)
	assertValidSendMessageResult(t, result)
}

func Test_Sending_RequestWithoutFromField_ResultBadRequest(t *testing.T) {

	result, ok := testPostNewMessage(t, Message{
		From:    "",
		Subject: "Test sending requestWithoutFromField result bad request",
		Content: "Hello?",
		To: []string{
			"fulano@mail.com",
		},
		Type: PlainMail,
	}, http.StatusBadRequest)

	assert.True(t, ok)
	assertErrorSendMessageResult(t, result, 1)
}

func Test_Sending_RequestWithoutToField_ResultBadRequest(t *testing.T) {

	result, ok := testPostNewMessage(t, Message{
		From:    "mail@mail.com",
		Subject: "Test sending requestWithoutFromField result bad request",
		Content: "Hello?",
		To:      []string{},
		Type:    PlainMail,
	}, http.StatusBadRequest)

	assert.True(t, ok)
	assertErrorSendMessageResult(t, result, 1)
}

func Test_Sending_NoMessage_ResultBadRequest(t *testing.T) {

	_, ok := testPostNewMessage(t, "", http.StatusBadRequest)
	assert.False(t, ok, "post no message (returns error)")
}

///////////////////////////////////

func createDefaultMailSender() MessageSender {
	return MailMessageSender{
		Config: config.SMTP,
	}
}

func assertErrorSendMessageResult(t *testing.T, result SendMessageResult, qtdErrors int) {
	assert.False(t, result.Success)
	assert.Equal(t, qtdErrors, len(result.Errors))
}

func assertValidSendMessageResult(t *testing.T, result SendMessageResult) {
	assert.True(t, result.Success)
	assert.True(t, result.MessageID != "")
	assert.True(t, len(result.Errors) == 0)
	assert.False(t, result.Date.IsZero())
}

func testPostNewMessage(t *testing.T, msg interface{}, expectedStatusCode int) (result SendMessageResult, ok bool) {

	ok = false
	result = SendMessageResult{}
	msgBytes, _ := json.Marshal(msg)
	buffer := bytes.NewBuffer(msgBytes)

	resp, err := client.Post(server.URL+"/messages", "application/json", buffer)

	if assert.Nil(t, err, "post new message request (no error)") {

		if assert.Equal(t, expectedStatusCode, resp.StatusCode) {
			defer resp.Body.Close()

			respBuffer := bytes.Buffer{}
			n, _ := respBuffer.ReadFrom(resp.Body)

			if assert.True(t, n > 0) {
				err := json.Unmarshal(respBuffer.Bytes(), &result)

				if err == nil {
					ok = true
				}
			}
		}
	}

	return
}
