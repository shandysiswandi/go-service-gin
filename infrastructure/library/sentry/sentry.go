package sentry

import (
	"go-service-gin/config"

	"github.com/getsentry/sentry-go"
)

// Sentry is
type Sentry struct {
	config *config.SentryConfig
}

// New is
func New(sc *config.SentryConfig) *Sentry {
	if sc == nil {
		return nil
	}

	if err := sentry.Init(sentry.ClientOptions{Dsn: sc.DSN, Environment: sc.ENV, Release: sc.Release}); err != nil {
		return nil
	}

	return &Sentry{sc}
}

// Error is
func (s *Sentry) Error(e error) *sentry.EventID {
	return sentry.CaptureException(e)
}

// Message is
func (s *Sentry) Message(msg string) *sentry.EventID {
	return sentry.CaptureMessage(msg)
}
