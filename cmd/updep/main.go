package main

import (
	"log/slog"
	"os"

	"updep/pkg/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "")
		if err != nil {
			slog.Error("error opening file for logging", "err", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(app.New())
	defer p.Kill()
	if _, err := p.Run(); err != nil {
		slog.Error("error running program", "err", err)
		os.Exit(1)
	}
}
