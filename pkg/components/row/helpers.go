package row

import (
	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"

	"github.com/charmbracelet/lipgloss"
)

func CalculateColumnWidths(
	packages []packagemanager.Package,
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
