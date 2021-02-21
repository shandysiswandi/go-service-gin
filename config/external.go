package config

import "os"

// ExternalConfig is
type ExternalConfig struct {
	JSONPlaceholder string
}

// NewExternalConfig is
func NewExternalConfig() *ExternalConfig {
	return &ExternalConfig{
		JSONPlaceholder: os.Getenv("URL_JSONPLACEHOLDER"),
	}
}
