package pkgtable

import (
	"updep/pkg/config"
	"updep/pkg/packages"

	"github.com/charmbracelet/lipgloss"
)

func calculateColumnWidths(
	packages []packages.Package,
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
	pkg packages.Package,
	columnWidths [config.ColumnCount]int,
) string {
	nameCellStyle := lipgloss.NewStyle()

	if pkg.Target != nil {
		switch *pkg.Target {
		case pkg.Wanted:
			nameCellStyle = minorDiffStyle
		case pkg.Latest:
			nameCellStyle = majorDiffStyle
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center,
		nameCellStyle.Width(columnWidths[0]+config.ColumnGap).Render(pkg.Name),
		lipgloss.PlaceHorizontal(
			columnWidths[1]+config.ColumnGap,
			lipgloss.Left,
			pkg.Wanted.RenderWithDiff(pkg.Current),
		),
		lipgloss.PlaceHorizontal(
			columnWidths[2]+config.ColumnGap,
			lipgloss.Left,
			pkg.Latest.RenderWithDiff(pkg.Current),
		),
		lipgloss.PlaceHorizontal(
			columnWidths[3],
			lipgloss.Left,
			pkg.Current.String(),
		),
	)
}
