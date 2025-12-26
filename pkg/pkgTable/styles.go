package pkgtable

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle    = lipgloss.NewStyle().Bold(true).Underline(true)
	activeRowStyle = lipgloss.NewStyle().
			Bold(true).
			Background(config.Theme.ActiveBackground)
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Major)
	minorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Minor)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Patch)
)
