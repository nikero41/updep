package depstable

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerContainerStyle = lipgloss.NewStyle().Margin(1, 0)
	headerStyle          = lipgloss.NewStyle().
				Bold(true).
				Background(config.Theme.Primary).
				Foreground(config.Theme.PrimaryText).
				Padding(0, 1)
	helpContainerStyle = lipgloss.NewStyle().Margin(1, 0, 0, 2)

	cursorStyle = lipgloss.NewStyle().
			Foreground(config.Theme.Primary).
			Bold(true)
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Major)
	minorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Minor)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Patch)
)
