package depstable

import (
	"fmt"

	"updep/pkg/config"
	"updep/pkg/dependency"
	"updep/pkg/version"

	"github.com/charmbracelet/lipgloss"
)

func renderRow(
	d dependency.Dependency,
	columnWidths [config.ColumnCount]int,
) string {
	nameCellStyle := lipgloss.NewStyle()

	switch d.DiffLevel() {
	case version.Major:
		nameCellStyle = majorDiffStyle
	case version.Minor:
		nameCellStyle = minorDiffStyle
	case version.Patch:
		nameCellStyle = patchDiffStyle
	}

	renderedWanted := renderWithDiff(
		d.Wanted,
		d.Current,
		d.Target == dependency.Wanted,
	)
	renderedLatest := renderWithDiff(
		d.Latest,
		d.Current,
		d.Target == dependency.Latest,
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		renderRowColumn(nameCellStyle.Render(d.Name), columnWidths[0], true),
		renderRowColumn(renderedWanted, columnWidths[1], true),
		renderRowColumn(renderedLatest, columnWidths[2], true),
		renderRowColumn(d.Current.String(), columnWidths[3], false),
	)
}

func renderRowColumn(label string, width int, withGap bool) string {
	if withGap {
		width += config.ColumnGap
	}
	return lipgloss.PlaceHorizontal(width, lipgloss.Left, label)
}

func renderWithDiff(v, target version.Version, selected bool) string {
	baseStyle := lipgloss.NewStyle().Underline(selected)

	switch version.VersionDiffLevel(v, target) {
	case version.Major:
		return majorDiffStyle.Underline(selected).Render(v.String())

	case version.Minor:
		return baseStyle.Render(fmt.Sprintf(
			"%d.",
			v.Major,
		)) +
			minorDiffStyle.Underline(selected).
				Render(fmt.Sprintf("%d.%d%s", v.Minor, v.Patch, v.Suffix))

	case version.Patch:
		return baseStyle.Render(fmt.Sprintf(
			"%d.%d.",
			v.Major,
			v.Minor,
		)) +
			patchDiffStyle.Underline(selected).
				Render(fmt.Sprintf("%d%s", v.Patch, v.Suffix))

	default:
		return v.String()
	}
}
