package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
)

const (
	defaultMsg = "msg"
)

type Option func(*ZapLogger)

type ZapLogger struct {
	trackingId string
	logger     *zap.Logger
	ctx        context.Context
	msg        string
}

func NewLogger(l *zap.Logger, opts ...Option) *ZapLogger {
	logger := &ZapLogger{
		logger: l,
		msg:    defaultMsg,
	}
	for _, o := range opts {
		o(logger)
	}
	return logger
}

func WithTracingId(tracingId string) Option {
	return func(logger *ZapLogger) {
		logger.trackingId = tracingId
	}
}

func WithContext(ctx context.Context) Option {
	return func(logger *ZapLogger) {
		ctx := context.WithValue(ctx, "logger", logger)
		logger.ctx = ctx
	}
}

func (l *ZapLogger) Log(level Level, kvs ...interface{}) error {
	if len(kvs) == 0 || len(kvs)%2 != 0 {
		l.logger.Warn(fmt.Sprint("Key-value must appear in pairs: ", kvs))
		kvs = append(kvs, "")
	}
	if l.trackingId != "" {
		kvs = append(kvs, "tracing_id", l.trackingId)
	}
	var fields []zap.Field
	for i := 0; i < len(kvs); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(kvs[i]), kvs[i+1]))
	}

	switch level {
	case Debug:
		l.logger.Debug("", fields...)
	case Info:
		l.logger.Info("", fields...)
	case Warn:
		l.logger.Warn("", fields...)
	case Error:
		l.logger.Error("", fields...)
	case Fatal:
		l.logger.Fatal("", fields...)
	}
	return nil
}

// Close - Freeze buffers
func (l *ZapLogger) Close() error {
	return l.logger.Sync()
}

// Debug - logger a message at debug level
func (l *ZapLogger) Debug(kvs ...interface{}) {
	_ = l.Log(Debug, l.msg, fmt.Sprint(kvs...))
}

// Debugw - logger a message at debug level
func (l *ZapLogger) Debugw(kvs ...interface{}) {
	_ = l.Log(Debug, kvs...)
}

// Debugf - logger a message at debug level
func (l *ZapLogger) Debugf(format string, kvs ...interface{}) {
	_ = l.Log(Debug, l.msg, fmt.Sprintf(format, kvs...))
}

// Info - logger a message at info level
func (l *ZapLogger) Info(kvs ...interface{}) {
	_ = l.Log(Info, l.msg, fmt.Sprint(kvs...))
}

// Infow - logger a message at info level
func (l *ZapLogger) Infow(kvs ...interface{}) {
	_ = l.Log(Info, kvs...)
}

// Infof - logger a message at info level
func (l *ZapLogger) Infof(format string, kvs ...interface{}) {
	_ = l.Log(Info, l.msg, fmt.Sprintf(format, kvs...))
}

// Warn - logger a message at warn level
func (l *ZapLogger) Warn(kvs ...interface{}) {
	_ = l.Log(Warn, l.msg, fmt.Sprint(kvs...))
}

// Warnw - logger a message at warn level
func (l *ZapLogger) Warnw(kvs ...interface{}) {
	_ = l.Log(Warn, kvs...)
}

// Warnf - logger a message at warn level
func (l *ZapLogger) Warnf(format string, kvs ...interface{}) {
	_ = l.Log(Warn, l.msg, fmt.Sprintf(format, kvs...))
}

// Error - logger a message at error level
func (l *ZapLogger) Error(kvs ...interface{}) {
	_ = l.Log(Error, l.msg, fmt.Sprint(kvs...))
}

// Errorw - logger a message at error level
func (l *ZapLogger) Errorw(kvs ...interface{}) {
	_ = l.Log(Error, kvs...)
}

// Errorf - logger a message at error level
func (l *ZapLogger) Errorf(format string, kvs ...interface{}) {
	_ = l.Log(Error, l.msg, fmt.Sprintf(format, kvs...))
}

// Fatal - logger a message at fatal level
func (l *ZapLogger) Fatal(kvs ...interface{}) {
	_ = l.Log(Fatal, l.msg, fmt.Sprint(kvs...))
	os.Exit(1)
}

// Fatalw - logger a message at fatal level
func (l *ZapLogger) Fatalw(kvs ...interface{}) {
	_ = l.Log(Fatal, kvs...)
	os.Exit(1)
}

// Fatalf - logger a message at fatal level
func (l *ZapLogger) Fatalf(format string, kvs ...interface{}) {
	_ = l.Log(Fatal, kvs...)
	os.Exit(1)
}
