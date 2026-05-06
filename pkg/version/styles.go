package version

import (
	"updep/pkg/config"

	"charm.land/lipgloss/v2"
)

var (
	majorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Major)
	minorDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Minor)
	patchDiffStyle = lipgloss.NewStyle().Foreground(config.Theme.Patch)
)
