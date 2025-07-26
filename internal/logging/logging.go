package logging

import (
	"log/slog"
	"os"
)

func NewLogger(logLvl string) (logger *slog.Logger) {
	logLevel := ConvertLogLevel(logLvl)
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	logger = slog.New(logHandler)
	return
}

func ConvertLogLevel(cfgLogLevel string) slog.Level {
	switch cfgLogLevel {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
