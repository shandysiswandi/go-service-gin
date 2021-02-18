package logger_test

import (
	"go-service-gin/util/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogInfo(t *testing.T) {
	assert.Nil(t, logger.LogInfo("info"))
}

func TestLogError(t *testing.T) {
	assert.Nil(t, logger.LogError("error"))
}
