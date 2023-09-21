package blablo

import (
	"log/slog"
	"os"

	pkg "github.com/oleksii-honchar/blablo/go-pkg"
)

type Logger struct {
	*slog.Logger // Embed the slog.Logger to extend it if needed
}

func NewLogger(prefix string) *Logger {
	opts := pkg.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := pkg.NewPrettyHandler(prefix, os.Stdout, opts)
	logger := slog.New(handler)

	return &Logger{Logger: logger}
}
