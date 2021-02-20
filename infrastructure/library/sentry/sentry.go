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

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         sc.DSN,
		Environment: sc.ENV,
	})
	if err != nil {
		return nil
	}
	return &Sentry{sc}
}

// CaptureError is
func (s *Sentry) CaptureError(e error) *sentry.EventID {
	return sentry.CaptureException(e)
}

// CaptureMessage is
func (s *Sentry) CaptureMessage(msg string) *sentry.EventID {
	return sentry.CaptureMessage(msg)
}
