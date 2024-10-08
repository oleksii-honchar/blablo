package goPkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"

	c "github.com/oleksii-honchar/coteco"
)

type PrettyHandlerOptions struct {
	SlogOpts  slog.HandlerOptions
	UseColors bool
}

type PrettyHandler struct {
	slog.Handler
	logger *log.Logger
}

func (handler *PrettyHandler) Handle(ctx context.Context, rec slog.Record) error {
	level := rec.Level.String()

	switch rec.Level {
	case slog.LevelDebug:
		level = c.WithMagenta(level)
	case slog.LevelInfo:
		level = c.WithBlue(level)
	case slog.LevelWarn:
		level = c.WithYellow(level)
	case slog.LevelError:
		level = c.WithRed(level)
	}
	level += " |"

	fields := make(map[string]interface{}, rec.NumAttrs())
	rec.Attrs(func(attr slog.Attr) bool {
		fields[attr.Key] = attr.Value.Any()

		return true
	})

	jsonStr, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	// timeStr := rec.Time.Format("2006/01/02 15:05:05.000")
	msg := c.White() + rec.Message + c.Reset()

	// line := fmt.Sprintf("%s %s %s", c.Gray247+timeStr+c.Reset, level, msg)
	line := fmt.Sprintf("%s %s", level, msg)

	if len(fields) > 0 {
		line += " " + c.Yellow() + string(jsonStr) + c.Reset()
	}

	handler.logger.Println(line)

	return nil
}

func NewPrettyHandler(
	prefix string,
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	c.UseColors = opts.UseColors

	normPrefix := ""
	if opts.SlogOpts.Level == slog.LevelDebug {
		normPrefix = TrimOrPadStringRight(prefix, 10) + " | "
	}

	handler := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		logger:  log.New(out, normPrefix, 0),
	}

	return handler
}
