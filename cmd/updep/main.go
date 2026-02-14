package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"updep/pkg/app"
	"updep/pkg/logger"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic", "err", err)
			os.Exit(1)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		slog.Error("error executing command", "err", err)
		os.Exit(1)
	}
}

var loggerConfig *logger.LoggerConfig

var rootCmd = &cobra.Command{
	Use: "updep",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		loggerConfig, err = logger.New()
		if err != nil {
			slog.Error("error creating logger", "err", err)
			return err
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		loggerConfig.Close()
	},
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(app.New())
		defer p.Kill()
		if _, err := p.Run(); err != nil {
			slog.Error("error running program", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
}
