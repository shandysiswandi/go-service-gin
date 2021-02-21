package config_test

import (
	"go-service-gin/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExternalConfig(t *testing.T) {
	is := assert.New(t)

	os.Setenv("URL_JSONPLACEHOLDER", "url")

	actual := config.NewExternalConfig()

	is.NotEqual(&config.ExternalConfig{}, actual)
	is.Equal("url", actual.JSONPlaceholder)
}
