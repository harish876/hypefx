package utils

import (
	"log/slog"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

func FromConfig(path string) *lumberjack.Logger {
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(path),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	return logger
}

func FromLogLevel(level string) slog.Level {
	formattedLevel := strings.ToUpper(level)
	switch formattedLevel {
	case "LEVEL":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
