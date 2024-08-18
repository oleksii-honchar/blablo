package blablo

import (
	"log/slog"
	"os"

	pkg "github.com/oleksii-honchar/blablo/go-pkg"
)

type Logger struct {
	*slog.Logger // Embed the slog.Logger to extend it if needed
}

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

func getLogLevelFromString(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func NewLogger(prefix string, logLevel string, useColors bool) *Logger {
	opts := pkg.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: getLogLevelFromString(logLevel),
		},
		UseColors: useColors,
	}
	handler := pkg.NewPrettyHandler(prefix, os.Stdout, opts)
	logger := slog.New(handler)

	return &Logger{Logger: logger}
}
