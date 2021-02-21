package sentry_test

import (
	"errors"
	"go-service-gin/config"
	"go-service-gin/infrastructure/library/sentry"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("SENTRY_DSN", "https://8e4695acf05e4f07b30fde8b22082d84@o528539.ingest.sentry.io/5645935")
	os.Setenv("SENTRY_RELEASE", "1")
	os.Setenv("SENTRY_ENVIRONMENT", "debug")
	os.Setenv("SENTRYGODEBUG", "1,2")
}

func TestSentry_Message_And_Error(t *testing.T) {
	is := assert.New(t)

	s := sentry.New(config.NewSentryConfig())
	msg := s.Message("message")
	exc := s.Error(errors.New("error"))
	//
	is.NotNil(s)
	is.NotNil(msg)
	is.NotNil(exc)
}
