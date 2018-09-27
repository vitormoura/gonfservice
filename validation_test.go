package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageWithoutFromField_ValidationFail(t *testing.T) {

	m := Message{
		From:    "",
		Content: "Um exemplo",
		Subject: "Test sending validRequest result OK",
		To: []string{
			"fulano@mail.com",
		},
	}

	errors := validateMailMessage(&m)

	assert.True(t, len(errors) == 1)
	assert.Equal(t, errors[0], MsgInvalidMessage)
}

func Test_MessageWithoutContent_ValidationFail(t *testing.T) {

	m := Message{
		From:    "chico@mail.com",
		Content: "",
		Subject: "Test sending validRequest result OK",
		To: []string{
			"fulano@mail.com",
		},
	}

	errors := validateMailMessage(&m)

	assert.True(t, len(errors) == 1)
}
