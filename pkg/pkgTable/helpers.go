package pkgtable

import (
	"fmt"

	"updep/pkg/config"
	"updep/pkg/dependency"
	"updep/pkg/version"

	"github.com/charmbracelet/lipgloss"
)

func calculateColumnWidths(
	packages []dependency.Dependency,
) [config.ColumnCount]int {
	columnWidths := [config.ColumnCount]int{}
	for _, p := range packages {
		columnWidths[0] = max(columnWidths[0], lipgloss.Width(p.Name))
		columnWidths[1] = max(columnWidths[1], lipgloss.Width(p.Wanted.String()))
		columnWidths[2] = max(columnWidths[2], lipgloss.Width(p.Latest.String()))
		columnWidths[3] = max(columnWidths[3], lipgloss.Width(p.Current.String()))
	}

	return columnWidths
}

func renderRow(
	pkg dependency.Dependency,
	columnWidths [config.ColumnCount]int,
) string {
	nameCellStyle := lipgloss.NewStyle()

	var diffLevel version.DiffLevel
	switch pkg.Target {
	case dependency.None:
		diffLevel = version.None
	case dependency.Wanted:
		diffLevel = version.VersionDiffLevel(pkg.Current, pkg.Wanted)
	case dependency.Latest:
		diffLevel = version.VersionDiffLevel(pkg.Current, pkg.Latest)
	}

	if pkg.Target != dependency.None {
		switch diffLevel {
		case version.Major:
			nameCellStyle = majorDiffStyle
		case version.Minor:
			nameCellStyle = minorDiffStyle
		case version.Patch:
			nameCellStyle = patchDiffStyle
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		renderRowColumn(nameCellStyle.Render(pkg.Name), columnWidths[0], true),
		renderRowColumn(
			renderWithDiff(
				pkg.Wanted,
				pkg.Current,
				pkg.Target == dependency.Wanted,
			),
			columnWidths[1],
			true,
		),
		renderRowColumn(
			renderWithDiff(
				pkg.Latest,
				pkg.Current,
				pkg.Target == dependency.Latest,
			),
			columnWidths[2],
			true,
		),
		renderRowColumn(pkg.Current.String(), columnWidths[3], false),
	)
}

func renderRowColumn(label string, width int, withGap bool) string {
	if withGap {
		width += config.ColumnGap
	}

	return lipgloss.PlaceHorizontal(width, lipgloss.Left, label)
}

func renderWithDiff(v, b version.Version, selected bool) string {
	baseStyle := lipgloss.NewStyle().Underline(selected)

	switch version.VersionDiffLevel(v, b) {
	case version.Major:
		return majorDiffStyle.Underline(selected).Render(v.String())

	case version.Minor:
		return baseStyle.Render(fmt.Sprintf(
			"%d.",
			v.Major,
		)) +
			minorDiffStyle.Underline(selected).
				Render(fmt.Sprintf("%d.%d", v.Minor, v.Patch))

	case version.Patch:
		return baseStyle.Render(fmt.Sprintf(
			"%d.%d.",
			v.Major,
			v.Minor,
		)) +
			patchDiffStyle.Underline(selected).Render(fmt.Sprintf("%d", v.Patch))

	default:
		return v.String()
	}
}
