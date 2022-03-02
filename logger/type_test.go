package logger

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

func TestProductionModeFile(t *testing.T) {
	fileLogger, err := NewFileLogger(ProductionMode, &lumberjack.Logger{
		Filename: "./2022-02-28.logger",
		MaxSize: 10,
		MaxBackups: 10,
		MaxAge: 20,
		Compress: true,
		LocalTime: true,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger := NewLogger(fileLogger, WithTracingId(uuid.New().String()))

	logger.Infof("logger %v", "test")
}


func TestProductionModeDefault(t *testing.T) {
	defaultLogger, err := NewDefaultLogger(DevelopmentMode)
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger := NewLogger(defaultLogger, WithTracingId(uuid.New().String()))
	logger.Info(fmt.Sprintf("shop_id: %v", 1))
	logger.Infof("shop_id %v", 1)
	logger.Infow("shop_id", 1, "user_id", 2)
}
