package config

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type colors struct {
	Primary          color.Color
	PrimaryText      color.Color
	Major            color.Color
	Minor            color.Color
	Patch            color.Color
	ActiveBackground color.Color
	Success          color.Color
}

var Theme = colors{
	Primary:          lipgloss.Color("#cba6f7"),
	PrimaryText:      lipgloss.Color("#1e1e2e"),
	Major:            lipgloss.Color("#f38ba8"),
	Minor:            lipgloss.Color("#fab387"),
	Patch:            lipgloss.Color("#a6e3a1"),
	ActiveBackground: lipgloss.Color("#45475a"),
	Success:          lipgloss.Color("#a6e3a1"),
}
