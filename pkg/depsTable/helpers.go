package depstable

import (
	"math"

	"updep/pkg/config"
	"updep/pkg/dependency"

	"github.com/charmbracelet/lipgloss"
)

func (t *DepsTable) scrollFix() {
	if t.cursor < t.offset {
		t.offset = t.cursor
	}
	headerHeight := lipgloss.Height(headerView([config.ColumnCount]int{}))
	helpHeight := lipgloss.Height(
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
	renderedRowsHeight := int(
		math.Max(
			math.Min(
				float64(t.Height-headerHeight-helpHeight),
				float64(len(t.dependencies)),
			),
			0),
	)

	if t.cursor > t.offset+renderedRowsHeight-1 {
		t.offset = t.cursor - renderedRowsHeight + 1
	}

	if t.offset > len(t.dependencies)-renderedRowsHeight {
		t.offset = len(t.dependencies) - renderedRowsHeight
	}
}

func (t DepsTable) rowCount() int {
	headerHeight := lipgloss.Height(headerView(t.columnWidths))
	helpHeight := lipgloss.Height(
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
	rowCount := int(
		math.Max(
			math.Min(
				float64(t.Height-headerHeight-helpHeight),
				float64(len(t.dependencies)),
			),
			0),
	)

	return int(
		math.Min(float64(rowCount), float64(len(t.dependencies))),
	)
}

func (t *DepsTable) SetDependencies(deps []dependency.Dependency) {
	t.dependencies = deps
	t.columnWidths = calculateColumnWidths(deps)
}

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
