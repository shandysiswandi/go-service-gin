package config_test

import (
	"go-service-gin/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppConfig(t *testing.T) {
	is := assert.New(t)

	os.Setenv("ENV", "testing")
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "8080")
	os.Setenv("TZ", "UTC")

	actual := config.NewAppConfig()

	is.NotEqual(&config.AppConfig{}, actual)
	is.Equal("testing", actual.Env)
	is.Equal("127.0.0.1", actual.Host)
	is.Equal("8080", actual.Port)
	is.Equal("UTC", actual.Timezone)
}
