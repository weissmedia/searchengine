package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()

	// Default log level: info
	logLevel := zapcore.InfoLevel

	// Check if the log level is set via an environment variable
	if level, exists := os.LookupEnv("SEARCHENGINE_LOG_LEVEL"); exists {
		if parsedLevel, err := zapcore.ParseLevel(level); err == nil {
			logLevel = parsedLevel
		}
	}

	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// GetLogger returns the configured zap logger
func GetLogger() *zap.Logger {
	return logger
}

// SyncLogger should be called at the end of the application to flush any buffered log entries.
func SyncLogger() {
	if err := logger.Sync(); err != nil {
		logger.Error("Error syncing logger", zap.Error(err))
	}
}
