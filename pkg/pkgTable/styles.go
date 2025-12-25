package pkgtable

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle    = lipgloss.NewStyle().Bold(true).Underline(true)
	activeRowStyle = lipgloss.NewStyle().
			Bold(true).
			Background(config.Theme.Surface1)
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Red)
	minorDiffStyle = lipgloss.NewStyle().
			Foreground(config.Theme.Peach)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Green)
)
