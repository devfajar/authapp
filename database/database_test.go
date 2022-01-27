package database

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDatabase()
	assert.NoError(t, err)
}
