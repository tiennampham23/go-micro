package logger

import (
	"context"
	"testing"
)

func InitLoggerInHttpInterceptor() context.Context {
	ctx := context.Background()
	logger := LogContext(ctx)
	return AddLoggerToCtx(ctx, logger)
}
func TestContextLogger(t *testing.T) {
	ctx := InitLoggerInHttpInterceptor()
	logger := LogContext(ctx)
	logger.Infof("Hello %v", "Nam")
	PassingContext(ctx)
}

func PassingContext(ctx context.Context) {
	logCtx := LogContext(ctx)
	logCtx.Infof("Hello again %v", "Nam")

}