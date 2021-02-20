package config

import "os"

// SentryConfig is
type SentryConfig struct {
	DSN string
	ENV string
}

// NewSentryConfig is
func NewSentryConfig() *SentryConfig {
	return &SentryConfig{
		DSN: os.Getenv("SENTRY_DSN"),
		ENV: os.Getenv("SENTRY_ENV"),
	}
}
