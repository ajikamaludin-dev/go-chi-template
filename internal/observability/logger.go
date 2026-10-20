package observability

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Channels struct {
	App   *zap.Logger
	Error *zap.Logger
	Audit *zap.Logger
}

func NewLogger(dir, level string, maxSizeMB, maxBackups, maxAgeDays int, service string) (*Channels, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	lvl := parseZapLevel(level)

	encCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	prefix := ""
	if service != "" {
		prefix = service + "-"
	}

	makeLogger := func(filename string) *zap.Logger {
		w := &lumberjack.Logger{
			Filename:   filepath.Join(dir, prefix+filename),
			MaxSize:    maxSizeMB,
			MaxBackups: maxBackups,
			MaxAge:     maxAgeDays,
			Compress:   true,
		}
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encCfg),
			zapcore.AddSync(w),
			lvl,
		)
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	return &Channels{
		App:   makeLogger("app.log"),
		Error: makeLogger("error.log"),
		Audit: makeLogger("audit.log"),
	}, nil
}

func parseZapLevel(s string) zapcore.Level {
	switch strings.ToLower(s) {
	case "debug":
		return zapcore.DebugLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
