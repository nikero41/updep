package version

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Red)
	minorDiffStyle = lipgloss.NewStyle().
			Foreground(config.Theme.Peach)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Green)
)
