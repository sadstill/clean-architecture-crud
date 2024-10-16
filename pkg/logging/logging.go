package logging

import (
	"log/slog"
	"os"
	"strings"
)

var logger *slog.Logger

func GetLogger() *slog.Logger {
	return logger
}

func SetupLogger(logLevel string) {
	level := parseLogLevel(logLevel)

	handler := slogpretty.NewPrettyHandler(os.Stdout)
	handler.SetLevel(level)

	logger = slog.New(handler)
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "fatal":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
