package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadValidConfigFile_ResultOK(t *testing.T) {
	c, err := loadConfig("app.config.toml")

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.SMTP.Host, "localhost")
	assert.Equal(t, c.SMTP.Port, 25)
}

func Test_LoadFileThatNotExists_ResultError(t *testing.T) {
	_, err := loadConfig("app.config.not_found.toml")

	assert.NotNil(t, err)
}
