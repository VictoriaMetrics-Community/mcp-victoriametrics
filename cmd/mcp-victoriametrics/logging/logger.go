package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

// Logger wraps slog.Logger with enabled flag for conditional logging
type Logger struct {
	*slog.Logger
	enabled bool
	file    *os.File
}

// Config holds logging configuration
type Config struct {
	Enabled bool
	Format  string // "text" or "json"
	Level   string // "debug", "info", "warn", "error"
	File    string // path to log file (empty = stderr)
}

// New creates a new Logger based on the provided configuration
func New(cfg Config) (*Logger, error) {
	if !cfg.Enabled {
		return &Logger{
			Logger:  slog.New(slog.NewTextHandler(io.Discard, nil)),
			enabled: false,
		}, nil
	}

	var output io.Writer = os.Stderr
	var file *os.File

	if cfg.File != "" {
		f, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		output = f
		file = f
	}

	level := parseLevel(cfg.Level)
	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(output, opts)
	} else {
		handler = slog.NewTextHandler(output, opts)
	}

	return &Logger{
		Logger:  slog.New(handler),
		enabled: true,
		file:    file,
	}, nil
}

// Close closes the log file if it was opened
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// IsEnabled returns true if logging is enabled
func (l *Logger) IsEnabled() bool {
	return l.enabled
}

// Fatal logs an error message and exits the program.
// Always outputs to stderr even when logging is disabled.
func (l *Logger) Fatal(msg string, args ...any) {
	if l.enabled {
		l.Error(msg, args...)
	} else {
		_, _ = fmt.Fprint(os.Stderr, "FATAL: "+msg)
		for i := 0; i < len(args); i += 2 {
			if i+1 < len(args) {
				_, _ = fmt.Fprintf(os.Stderr, " %v=%v", args[i], args[i+1])
			}
		}
		_, _ = fmt.Fprintln(os.Stderr)
	}
	os.Exit(1)
}

// InfoOrPrint logs with structured format if enabled, otherwise prints plain message to stderr.
// Use this for important messages that should always be visible (like startup info).
func (l *Logger) InfoOrPrint(msg string, args ...any) {
	if l.enabled {
		l.Info(msg, args...)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, msg)
	}
}

// parseLevel converts string level to slog.Level
func parseLevel(s string) slog.Level {
	switch s {
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
