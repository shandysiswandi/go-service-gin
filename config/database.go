package config

import "os"

// DatabaseConfig is
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Timezone string
}

// NewDatabaseConfig is
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}
}
