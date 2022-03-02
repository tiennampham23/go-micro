package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TypeLogger - type for logging: stdout or file
type TypeLogger string

// TypeLoggerMode - mode for logging: development or production
type TypeLoggerMode string

const (
	// FileType - file for logging
	FileType TypeLogger = "file"
	// DefaultType - default for logging
	DefaultType TypeLogger = "default"
	// DevelopmentMode - support logger in development mode
	DevelopmentMode TypeLoggerMode = "development"
	// ProductionMode - support logger in production mode
	ProductionMode TypeLoggerMode = "production"
)

func NewFileLogger(mode TypeLoggerMode, config *lumberjack.Logger) (*zap.Logger, error) {
	writeSync := zapcore.AddSync(config)
	var encoderConfig zapcore.EncoderConfig
	if mode == DevelopmentMode {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else if mode == ProductionMode {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		panic(fmt.Sprintf("Not support this mode %v", mode))
	}
	// Set default encode time
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Get atomic level
	atomicLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	// Get core
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSync, atomicLevel)
	zapcore.NewTee()
	logger := zap.New(core)
	logger = logger.WithOptions(
		zap.AddCallerSkip(2),                  // increases the number of callers skipped by caller annotation
		zap.AddStacktrace(zapcore.ErrorLevel), // record a stack trace for all error messages
	)
	return logger, nil
}

func NewDefaultLogger(mode TypeLoggerMode) (*zap.Logger, error) {
	var encoderConfig zapcore.EncoderConfig
	var cfg zap.Config
	if mode == DevelopmentMode {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg = zap.NewDevelopmentConfig()
	} else if mode == ProductionMode {
		encoderConfig = zap.NewProductionEncoderConfig()
		cfg = zap.NewProductionConfig()
	} else {
		panic(fmt.Sprintf("Not support this mode %v", mode))
	}
	// Set default encode time
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig = encoderConfig
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	logger = logger.WithOptions(
		zap.AddCallerSkip(2),                  // increases the number of callers skipped by caller annotation
		zap.AddStacktrace(zapcore.ErrorLevel), // record a stack trace for all error messages
	)
	return logger, nil
}
