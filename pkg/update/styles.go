package update

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	checkMarkStyle = lipgloss.NewStyle().Foreground(config.Theme.Success)
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Major)
	minorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Minor)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Patch)
)
