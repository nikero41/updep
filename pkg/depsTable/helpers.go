package depstable

import (
	"fmt"

	"updep/pkg/config"
	"updep/pkg/dependency"
	"updep/pkg/version"

	"github.com/charmbracelet/lipgloss"
)

func calculateColumnWidths(
	deps []dependency.Dependency,
) [config.ColumnCount]int {
	columnWidths := [config.ColumnCount]int{}
	for _, d := range deps {
		columnWidths[0] = max(columnWidths[0], lipgloss.Width(d.Name))
		columnWidths[1] = max(columnWidths[1], lipgloss.Width(d.Wanted.String()))
		columnWidths[2] = max(columnWidths[2], lipgloss.Width(d.Latest.String()))
		columnWidths[3] = max(columnWidths[3], lipgloss.Width(d.Current.String()))
	}

	return columnWidths
}

func renderRow(
	d dependency.Dependency,
	columnWidths [config.ColumnCount]int,
) string {
	nameCellStyle := lipgloss.NewStyle()

	var diffLevel version.DiffLevel
	switch d.Target {
	case dependency.None:
		diffLevel = version.None
	case dependency.Wanted:
		diffLevel = version.VersionDiffLevel(d.Current, d.Wanted)
	case dependency.Latest:
		diffLevel = version.VersionDiffLevel(d.Current, d.Latest)
	}

	if d.Target != dependency.None {
		switch diffLevel {
		case version.Major:
			nameCellStyle = majorDiffStyle
		case version.Minor:
			nameCellStyle = minorDiffStyle
		case version.Patch:
			nameCellStyle = patchDiffStyle
		}
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
			patchDiffStyle.Underline(selected).Render(fmt.Sprintf("%d%s", v.Patch, v.Suffix))

	default:
		return v.String()
	}
}
