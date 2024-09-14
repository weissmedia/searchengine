package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var logger *zap.Logger

// init initializes the global logger configuration. It sets the default log level to 'info'
// but allows customization via the SEARCHENGINE_LOG_LEVEL environment variable.
// The log format uses ISO8601 time encoding for consistent time formatting.
func init() {
	config := zap.NewProductionConfig()

	// Default log level: info
	logLevel := zapcore.InfoLevel

	// Check if the log level is set via an environment variable
	if level, exists := os.LookupEnv("SEARCHENGINE_LOG_LEVEL"); exists {
		// Parse and apply the log level from the environment variable if valid
		if parsedLevel, err := zapcore.ParseLevel(level); err == nil {
			logLevel = parsedLevel
		}
	}

	// Set the parsed or default log level
	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Build the logger
	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err) // Terminate the program if logger creation fails
	}
}

// GetLogger returns the globally configured zap logger instance.
// This allows the logger to be reused throughout the application.
func GetLogger() *zap.Logger {
	return logger
}

// SyncLogger should be called at the end of the application to flush any buffered log entries.
func SyncLogger() {
	if err := logger.Sync(); err != nil && !isExpectedSyncError(err) {
		logger.Error("Error syncing logger", zap.Error(err))
	}
}

// Helper function to ignore specific sync errors
func isExpectedSyncError(err error) bool {
	// Extend this check to cover more expected errors during testing
	return strings.Contains(err.Error(), "inappropriate ioctl for device") ||
		strings.Contains(err.Error(), "bad file descriptor")
}
