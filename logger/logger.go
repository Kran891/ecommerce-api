package logger

import (
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func Init() {
	var err error
	env := os.Getenv("ENV")
	if env == "PROD" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("❌❌❌ failed to initialize logger")
	}
	defer logger.Sync() // flushes buffer, if any
	logger.Info("✅✅✅ Logger initialized successfully")
}

// LogInfo logs an info message with optional key-value pairs
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// LogError logs an error message
func Error(msg string, err error, fields ...zap.Field) {
	logger.Error(msg, append(fields, zap.Error(err))...)
}

// LogDebug logs debug info
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// LogWarn logs warnings
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Sync flushes any buffered log entries
func SyncLogger() {
	_ = logger.Sync()
}
