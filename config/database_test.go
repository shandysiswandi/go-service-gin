package config_test

import (
	"go-service-gin/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseConfig_Mysql(t *testing.T) {
	is := assert.New(t)

	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual := config.NewDatabaseConfig()

	is.NotEqual(&config.DatabaseConfig{}, actual)
	is.Equal("mysql", actual.Driver)
	is.Equal("127.0.0.1", actual.Host)
	is.Equal("3306", actual.Port)
	is.Equal("root", actual.Username)
	is.Equal("", actual.Password)
	is.Equal("go-service", actual.Name)
	is.Equal("UTC", actual.Timezone)
}

func TestNewDatabaseConfig_Postgres(t *testing.T) {
	is := assert.New(t)

	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual := config.NewDatabaseConfig()

	is.NotEqual(&config.DatabaseConfig{}, actual)
	is.Equal("postgres", actual.Driver)
	is.Equal("127.0.0.1", actual.Host)
	is.Equal("5432", actual.Port)
	is.Equal("root", actual.Username)
	is.Equal("", actual.Password)
	is.Equal("go-service", actual.Name)
	is.Equal("UTC", actual.Timezone)
}

func TestNewDatabaseConfig_Mongo(t *testing.T) {
	is := assert.New(t)

	os.Setenv("DB_DRIVER", "mongo")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual := config.NewDatabaseConfig()

	is.NotEqual(&config.DatabaseConfig{}, actual)
	is.Equal("mongo", actual.Driver)
	is.Equal("127.0.0.1", actual.Host)
	is.Equal("27017", actual.Port)
	is.Equal("root", actual.Username)
	is.Equal("", actual.Password)
	is.Equal("go-service", actual.Name)
	is.Equal("UTC", actual.Timezone)
}
