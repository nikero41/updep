package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/charmbracelet/fang"

	"updep/pkg/logger"
)

var loggerConfig *logger.LoggerConfig

func main() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic", "err", err)
			os.Exit(1)
		}
	}()

	var err error
	loggerConfig, err = logger.New()
	if err != nil {
		slog.Error("error creating logger", "err", err)
		os.Exit(1)
	}
	defer loggerConfig.Close()

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		slog.Error("error executing command", "err", err)
		os.Exit(1)
	}
}
