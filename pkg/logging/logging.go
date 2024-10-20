package logging

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		_ = SetupLogger(Info, Text)
	})
	return logger
}

func SetupLogger(logLevel LogLevel, logFormat LogFormat) error {
	slogLevel, err := parseLogLevel(logLevel)
	if err != nil {
		return err
	}

	var handler slog.Handler

	switch logFormat {
	case Json:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel,
		})
	case Text:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel,
		})
	default:
		return fmt.Errorf("unsupported log format provided")
	}

	logger = slog.New(handler)

	return nil
}

func parseLogLevel(level LogLevel) (slog.Level, error) {
	switch level {
	case Debug:
		return slog.LevelDebug, nil
	case Info:
		return slog.LevelInfo, nil
	case Warn:
		return slog.LevelWarn, nil
	case Error:
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unsupported log level format")
	}
}
