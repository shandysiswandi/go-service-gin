package config

import "os"

// AppConfig is
type AppConfig struct {
	Env      string
	Host     string
	Port     string
	Timezone string
}

// NewAppConfig is
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Env:      os.Getenv("ENV"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Timezone: os.Getenv("TZ"),
	}
}
