package version

import (
	"updep/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Major)
	minorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Minor)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Patch)
)
