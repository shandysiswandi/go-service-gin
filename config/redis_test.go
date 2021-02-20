package config_test

import (
	"go-service-gin/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedisConfig(t *testing.T) {
	is := assert.New(t)

	os.Setenv("REDIS_ADDRESS", "localhost:6379")
	os.Setenv("REDIS_PASSWORD", "")
	os.Setenv("REDIS_DATABASE", "0")

	actual := config.NewRedisConfig()

	is.NotEqual(&config.RedisConfig{}, actual)
	is.Equal("localhost:6379", actual.Address)
	is.Equal("", actual.Password)
	is.Equal(0, actual.Database)

	os.Clearenv()

	os.Setenv("REDIS_DATABASE", "A")
	actual2 := config.NewRedisConfig()
	is.Equal(0, actual2.Database)
}
