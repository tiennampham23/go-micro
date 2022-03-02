package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger(zap.NewExample(), WithTracingId("xyz"))
	defer func() {
		_ = logger.Close()
	}()
	logger.Infof("logger %v", "debug")
}