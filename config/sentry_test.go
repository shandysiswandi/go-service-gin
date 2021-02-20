package config_test

import (
	"go-service-gin/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSentryConfig_NotNil(t *testing.T) {
	is := assert.New(t)

	os.Setenv("SENTRY_DSN", "url_sentry")
	os.Setenv("SENTRY_ENV", "testing")

	actual := config.NewSentryConfig()

	is.NotEqual(&config.SentryConfig{}, actual)
	is.Equal("url_sentry", actual.DSN)
	is.Equal("testing", actual.ENV)
}
