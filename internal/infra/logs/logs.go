package logs

import (
	"log/slog"
	"os"
)

type Logger struct {
	Logger *slog.Logger
}

func NewLogger() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelError)

	return &Logger{Logger: logger}
}
