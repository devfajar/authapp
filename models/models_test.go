package models

import (
	"authapp/database"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *Testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)

}
