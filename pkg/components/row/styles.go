package row

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	ActiveRowStyle = lipgloss.NewStyle().
			Bold(true).
			Background(config.Theme.Surface1)
	errorStyle = lipgloss.NewStyle().Foreground(config.Theme.Red)
	warningStyle       = lipgloss.NewStyle().
				Foreground(config.Theme.Peach)
	successStyle = lipgloss.NewStyle().Foreground(config.Theme.Green)

	optionalUpdateStyle = lipgloss.NewStyle().
				Foreground(config.Theme.Yellow)
)
