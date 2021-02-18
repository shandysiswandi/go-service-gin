package logger

import (
	log "github.com/sirupsen/logrus"
)

// LogInfo is
func LogInfo(args ...interface{}) error {
	log.Info(args...)
	return nil
}

// LogError is
func LogError(args ...interface{}) error {
	log.Error(args...)
	return nil
}
