package depstable

import (
	"updep/pkg/config"
	"updep/pkg/dependency"

	"charm.land/lipgloss/v2"
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
		t.cursor = max(t.offset+t.rowCount()-1, 0)
	} else if t.cursor < t.offset {
		t.cursor = t.offset
	}
}

func (t *DepsTable) scrollFix() {
	if t.cursor < t.offset {
		t.offset = t.cursor
		return
	}
	rowCount := t.rowCount()

	if t.cursor > t.offset+rowCount-1 {
		t.offset = max(t.cursor-rowCount+1, 0)
	}

	if t.offset > len(t.dependencies)-rowCount {
		t.offset = max(len(t.dependencies)-rowCount, 0)
	}
}

func (t DepsTable) rowCount() int {
	headerHeight := lipgloss.Height(t.header)
	helpHeight := lipgloss.Height(
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)

	return max(
		min(t.height-headerHeight-helpHeight, len(t.dependencies)),
		0)
}

func (t *DepsTable) SetDependencies(deps []dependency.Dependency) {
	t.dependencies = deps
	t.setColumnWidths(calculateColumnWidths(deps))
	t.cursor = 0
	t.scrollFix()
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
