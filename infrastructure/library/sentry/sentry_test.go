package sentry_test

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("SENTRY_DSN", "http")
	os.Setenv("SENTRY_RELEASE", "1")
	os.Setenv("SENTRY_ENVIRONMENT", "debug")
	os.Setenv("SENTRYGODEBUG", "1,2")
}

func TestSentry_CaptureMessage_And_CaptureError(t *testing.T) {
	// is := assert.New(t)

	// s := sentry.New(config.New())
	// msg := s.CaptureMessage("message")
	// exc := s.CaptureError(errors.New("error"))

	// is.NotNil(s)
	// is.NotNil(msg)
	// is.NotNil(exc)
}
