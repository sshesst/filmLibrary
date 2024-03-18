package logging

import (
	"context"
	"log/slog"
	"os"
)

const (
	defaultLevel      = LevelInfo
	defaultAddSource  = false
	defaultIsJSON     = false
	defaultSetDefault = true
)

func NewLogger(opts ...LoggerOption) *Logger {
	config := &LoggerOptions{
		Level:      defaultLevel,
		AddSource:  defaultAddSource,
		IsJSON:     defaultIsJSON,
		SetDefault: defaultSetDefault,
	}

	for _, opt := range opts {
		opt(config)
	}

	options := &HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	var h Handler = NewTextHandler(os.Stdout, options)

	if config.IsJSON {
		h = NewJSONHandler(os.Stdout, options)
	}

	logger := New(h)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

type LoggerOptions struct {
	Level      Level
	AddSource  bool
	IsJSON     bool
	SetDefault bool
}

type LoggerOption func(*LoggerOptions)

func WithLevel(level string) LoggerOption {
	return func(o *LoggerOptions) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		o.Level = l
	}
}

func WithAddSource(addSource bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.AddSource = addSource
	}
}

func WithIsJSON(isJSON bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.IsJSON = isJSON
	}
}

func WithSetDefault(setDefault bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.SetDefault = setDefault
	}
}

func WithAttrs(ctx context.Context, attrs ...Attr) *Logger {
	logger := L(ctx)
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

func WithDefaultAttrs(logger *Logger, attrs ...Attr) *Logger {
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

func L(ctx context.Context) *Logger {
	return loggerFromContext(ctx)
}

func Default() *Logger {
	return slog.Default()
}
