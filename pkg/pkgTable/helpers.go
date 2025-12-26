package pkgtable

import (
	"updep/pkg/config"
	"updep/pkg/packages"

	"github.com/charmbracelet/lipgloss"
)

// calculateColumnWidths computes the maximum lipgloss display width for each of
// the four table columns (Name, Wanted, Latest, Current) across the provided
// packages. The returned array contains widths in that column order.
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

// renderRow renders a single package as a four-column lipgloss-formatted row.
// The name column is styled based on pkg.Target (wanted → minorDiffStyle, latest → majorDiffStyle).
// The other columns show Wanted and Latest rendered with diffs against Current and the Current value itself,
// using widths from columnWidths (and config.ColumnGap where applied).
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