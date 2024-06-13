package utils

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/harish876/hypefx/internal/cli/commands"
)

type Logger struct {
	*slog.Logger
}

var (
	DEFAULT_LOG_FILE  = "hypefxlog.log"
	DEFAULT_LOG_LEVEL = "INFO"
)

func NewLogger(defaultPath string) (*Logger, error) {
	logFile, err := commands.GetConfig("logfile")
	if err != nil && logFile == nil {
		//fmt.Printf("failed to find logfile from settings. Using Default: %s\n", DEFAULT_LOG_FILE)
		logFile = defaultPath
	}
	level, err := commands.GetConfig("level")
	if err != nil && level == nil {
		//fmt.Printf("failed to find loglevel from settings. Using Default: %s\n", DEFAULT_LOG_FILE)
		level = DEFAULT_LOG_LEVEL
	}
	absPath, err := filepath.Abs(logFile.(string))
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(absPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     getLogLevel(level.(string)),
	}))

	return &Logger{logger}, nil
}

func getLogLevel(level string) slog.Leveler {
	level = strings.ToUpper(level)
	switch level {
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
