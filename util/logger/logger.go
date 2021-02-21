package logger

import (
	log "github.com/sirupsen/logrus"
)

// Info is
func Info(args ...interface{}) error {
	log.Info(args...)
	return nil
}

// Error is
func Error(args ...interface{}) error {
	log.Error(args...)
	return nil
}
