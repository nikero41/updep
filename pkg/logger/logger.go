package logger

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/lmittmann/tint"
)

type LoggerConfig struct {
	*slog.Logger
	File  *os.File
	Level *slog.LevelVar
}

func New() (*LoggerConfig, error) {
	filePath := path.Join(
		os.ExpandEnv("$HOME"),
		".local",
		"share",
		"updep",
		"updep.log",
	)
	f, err := getOrCreateDataFile(filePath)
	if err != nil {
		return nil, err
	}

	level := &slog.LevelVar{}
	level.Set(slog.LevelInfo)

	if os.Getenv("DEBUG") == "1" {
		level.Set(slog.LevelDebug)
	}

	logger := slog.New(
		tint.NewHandler(f, &tint.Options{
			Level:      level,
			TimeFormat: time.DateTime,
		}),
	)
	slog.SetDefault(logger)
	slog.Debug("logger initialized", "path", filePath)

	return &LoggerConfig{
		Logger: logger,
		Level:  level,
		File:   f,
	}, nil
}

func (l LoggerConfig) SetLevel(level slog.Level) {
	l.Level.Set(level)
}

func (l LoggerConfig) Close() error {
	slog.Debug("logger closing")
	return l.File.Close()
}

func getOrCreateDataFile(filePath string) (*os.File, error) {
	_, err := os.Stat(path.Dir(filePath))

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		err = os.MkdirAll(pathError.Path, 0o755)
	}
	if err != nil {
		return nil, err
	}

	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
}
