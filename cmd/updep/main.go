package main

import (
	"log/slog"
	"os"

	"updep/pkg/app"
	"updep/pkg/config"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f := config.GetLogFile()
	defer f.Close()
	config.SetupLogger(f)

	p := tea.NewProgram(app.New())
	defer p.Kill()
	if _, err := p.Run(); err != nil {
		slog.Error("error running program", "err", err)
		os.Exit(1)
	}
}
