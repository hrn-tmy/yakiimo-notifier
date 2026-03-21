package logger

import (
	"log/slog"
	"os"
)

// NewLogger はログの初期化を行います
func NewLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}
