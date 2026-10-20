package observability

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Channels struct {
	App   *slog.Logger
	Error *slog.Logger
	Audit *slog.Logger
}

func NewLogger(dir, level string, maxSizeMB, maxBackups, maxAgeDays int) (*Channels, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	lvl := parseLevel(level)

	makeLogger := func(filename string) *slog.Logger {
		w := &lumberjack.Logger{
			Filename:   filepath.Join(dir, filename),
			MaxSize:    maxSizeMB,
			MaxBackups: maxBackups,
			MaxAge:     maxAgeDays,
			Compress:   true,
		}
		// tee to stdout in dev-friendly manner
		h := slog.NewJSONHandler(io.MultiWriter(os.Stdout, w), &slog.HandlerOptions{Level: lvl})
		return slog.New(h)
	}

	return &Channels{
		App:   makeLogger("app.log"),
		Error: makeLogger("error.log"),
		Audit: makeLogger("audit.log"),
	}, nil
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Err wraps error into slog attribute
func Err(err error) slog.Attr {
	if err == nil {
		return slog.String("error", "")
	}
	return slog.String("error", err.Error())
}

var _ = fmt.Sprintf
