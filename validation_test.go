package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageWithoutFromField_ValidationFail(t *testing.T) {

	m := Message{
		Type: PlainMail,
		From:    "",
		Content: "Um exemplo",
		Subject: "Test sending validRequest result OK",
		To: []string{
			"fulano@mail.com",
		},
	}

	errors := validateMailMessage(&m)

	assert.True(t, len(errors) == 1)
	assert.Equal(t, errors[0], "invalid email address: ''")
}

func Test_MessageNotSupportedType_ValidationFail(t *testing.T) {

	m := Message{
		From:    "chico@mail.com",
		Content: "Um exemplo",
		Subject: "Test sending validRequest result OK",
		To: []string{
			"fulano@mail.com",
		},
		Type: MessageType(999),
	}

	errors := validateMailMessage(&m)

	assert.Equal(t, len(errors), 1, "not supported type")
}

func Test_MessageWithoutContent_ValidationFail(t *testing.T) {

	m := Message{
		Type: PlainMail,
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

func Test_EmailListAllAddressesValid_ValidationOK(t *testing.T) {
	errors := []string{}
	emails := []string{
		"i1779101@mai.com",
		"fulano@mail.com.br",
		"beltrano@mail",
	}

	validateMailAddressList(emails, &errors)

	assert.Equal(t, 0, len(errors), "all email addresses are valid")
}

func Test_EmailListAllAddressesInvalid_ValidationFail(t *testing.T) {
	errors := []string{}
	emails := []string{
		"i1779101",
		"fulano@",
		" ",
		"    ",
		"@",
		"mail@mail..",
		"mail@kakak.com@mam",
		"mail@",
	}

	validateMailAddressList(emails, &errors)

	assert.Equal(t, len(emails), len(errors), "every email is invalid")
}
