package config

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

func GetLogFile() *os.File {
	f, err := os.OpenFile(
		os.ExpandEnv("$HOME/.local/share/updep/updep.log"),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0o600,
	)
	if err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			slog.Warn("creating log file", "path", pathError.Path)

			if err = os.MkdirAll(
				filepath.Dir(pathError.Path),
				0o700,
			); err != nil {
				slog.Error("error creating directory for logging", "error", err)
				os.Exit(1)
			}

			f, err = os.Create(pathError.Path)
			if err != nil {
				slog.Error("error creating file for logging", "error", err)
				os.Exit(1)
			}
			slog.Debug("log file created", "path", pathError.Path)
		} else {
			slog.Error("could not setup logging", "error", err)
			os.Exit(1)
		}
	}

	return f
}

func SetupLogger(f *os.File) {
	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.New(slog.NewTextHandler(f, opts))
	slog.SetDefault(handler)

	if os.Getenv("DEBUG") != "" {
		logLevel.Set(slog.LevelDebug)
	}
}
