package logger

import (
	"context"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	FileName  = "./logger/app.logger"
	MaxSize   = 10
	MaxBackup = 10
	MaxAge    = 20
)

func getLoggerFileConfig() *lumberjack.Logger {
	fileName := viper.GetString("logger.path_file")
	maxSize := viper.GetInt("logger.max_size")
	maxBackup := viper.GetInt("logger.max_back_up")
	maxAge := viper.GetInt("logger.max_age")
	if fileName == "" {
		fileName = FileName
	}
	if maxSize == 0 {
		maxSize = MaxSize
	}
	if maxBackup == 0 {
		maxBackup = MaxBackup
	}
	if maxAge == 0 {
		maxAge = MaxAge
	}
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxBackups: maxBackup,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		Compress:   true,
		LocalTime:  true,
	}
}

// NewDevelopmentFileLogger - init file logger with development mode
func NewDevelopmentFileLogger(ctx context.Context) (*ZapLogger, error) {
	cfg := getLoggerFileConfig()
	fileLogger, err := NewFileLogger(DevelopmentMode, cfg)
	if err != nil {
		return nil, err
	}
	return NewLogger(fileLogger, WithTracingId(uuid.New().String()), WithContext(ctx)), nil
}

// NewProductionFileLogger - init file logger with production mode
func NewProductionFileLogger(ctx context.Context) (*ZapLogger, error) {
	cfg := getLoggerFileConfig()
	fileLogger, err := NewFileLogger(ProductionMode, cfg)
	if err != nil {
		return nil, err
	}
	return NewLogger(fileLogger, WithTracingId(uuid.New().String()), WithContext(ctx)), nil
}

// NewDevelopmentDefaultLogger - init default logger with development mode
func NewDevelopmentDefaultLogger(ctx context.Context) (*ZapLogger, error) {
	var trackingId string
	if v, ok := ctx.Value("tracing_id").(string); ok {
		trackingId = v
	} else {
		trackingId = uuid.NewString()
	}
	if l, ok := ctx.Value("logger").(*ZapLogger); ok {
		return l, nil
	}
	defaultLogger, err := NewDefaultLogger(DevelopmentMode)
	if err != nil {
		return nil, err
	}
	logger := NewLogger(defaultLogger, WithContext(ctx), WithTracingId(trackingId))
	return logger, nil
}

// NewProductionDefaultLogger - init default logger with production mode
func NewProductionDefaultLogger(ctx context.Context) (*ZapLogger, error) {
	defaultLogger, err := NewDefaultLogger(ProductionMode)
	if err != nil {
		return nil, err
	}
	return NewLogger(defaultLogger, WithTracingId(uuid.New().String()), WithContext(ctx)), nil
}


func AddLoggerToCtx(ctx context.Context, logger Logger) context.Context {
	if _, ok := ctx.Value("logger").(*Logger); ok {
		return ctx
	}
	ctx = context.WithValue(ctx, "logger", logger)
	return ctx
}

func LogContext(ctx context.Context) *ZapLogger {
	if l, ok := ctx.Value("logger").(*ZapLogger); ok {
		return l
	}
	envMode := viper.GetString("env")
	logType := viper.GetString("log_type")
	if TypeLoggerMode(envMode) == DevelopmentMode {
		if TypeLogger(logType) == FileType {
			l, _ := NewDevelopmentFileLogger(ctx)
			return l
		} else if TypeLogger(logType) == DefaultType {
			l, _ := NewDevelopmentDefaultLogger(ctx)
			return l
		}
	} else if TypeLoggerMode(envMode) == ProductionMode {
		if TypeLogger(logType) == FileType {
			l, _ := NewProductionFileLogger(ctx)
			return l
		} else if TypeLogger(logType) == DefaultType {
			l, _ := NewProductionDefaultLogger(ctx)
			return l
		}
	}
	// Default log with development mode and default
	l, _ := NewDevelopmentDefaultLogger(ctx)
	return l
}