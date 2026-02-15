package depstable

import (
	"updep/pkg/config"
	"updep/pkg/dependency"

	"github.com/charmbracelet/lipgloss"
)

func (t *DepsTable) SetHeight(height int) {
	t.height = height
	t.scrollFix()
}

func (t *DepsTable) scrollTo(to int) {
	if to < 0 {
		to = 0
	} else if to >= len(t.dependencies)-t.rowCount() {
		to = max(len(t.dependencies)-t.rowCount(), 0)
	}

	t.offset = to

	if t.cursor > t.offset+t.rowCount()-1 {
		t.cursor = max(t.offset + t.rowCount() - 1)
	} else if t.cursor < t.offset {
		t.cursor = t.offset
	}
}

func (t *DepsTable) scrollFix() {
	if t.cursor < t.offset {
		t.offset = t.cursor
	}
	headerHeight := lipgloss.Height(headerView([config.ColumnCount]int{}))
	helpHeight := lipgloss.Height(
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
	renderedRowsHeight := max(
		min(
			t.height-headerHeight-helpHeight,
			len(t.dependencies),
		),
		0)

	if t.cursor > t.offset+renderedRowsHeight-1 {
		t.offset = t.cursor - renderedRowsHeight + 1
	}

	if t.offset > len(t.dependencies)-renderedRowsHeight {
		t.offset = len(t.dependencies) - renderedRowsHeight
	}
}

func (t DepsTable) rowCount() int {
	headerHeight := lipgloss.Height(t.header)
	helpHeight := lipgloss.Height(
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
	rowCount := max(
		min(
			t.height-headerHeight-helpHeight,
			len(t.dependencies),
		),
		0)

	return min(rowCount, len(t.dependencies))
}

func (t *DepsTable) SetDependencies(deps []dependency.Dependency) {
	t.dependencies = deps
	t.setColumnWidths(calculateColumnWidths(deps))
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

func (t *DepsTable) setColumnWidths(widths [config.ColumnCount]int) {
	t.columnWidths = widths
	t.header = headerView(t.columnWidths)
}
