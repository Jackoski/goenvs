package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad_Correct(t *testing.T) {
	err := Load()

	expected := "value"

	assert.EqualError(t, err, "env WRONGenv, do not contain separator `=`")
	assert.Equal(t, expected, os.Getenv("KEY"))
}
