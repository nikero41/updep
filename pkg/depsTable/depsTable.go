package depstable

import (
	"updep/pkg/config"
	"updep/pkg/dependency"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DepsTable struct {
	helpModel    help.Model
	cursor       int
	offset       int
	height       int
	dependencies []dependency.Dependency
	header       string
	columnWidths [config.ColumnCount]int
}

func New() *DepsTable {
	return &DepsTable{
		helpModel:    help.New(),
		cursor:       0,
		offset:       0,
		height:       0,
		dependencies: []dependency.Dependency{},
		header:       headerView([config.ColumnCount]int{}),
		columnWidths: [config.ColumnCount]int{},
	}
}

func (t DepsTable) Init() tea.Cmd {
	return nil
}

func (t DepsTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, t.handleKeyPress(msg))
	}

	return t, tea.Batch(cmds...)
}

func (t DepsTable) View() string {
	rowCount := t.rowCount()
	renderedRows := make([]string, rowCount)
	for i, d := range t.dependencies[t.offset:rowCount] {
		row := renderRow(d, t.columnWidths)
		renderedRows[i] = cursorView(i+t.offset == t.cursor) + row
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		t.header,
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
}

func headerView(columnWidths [config.ColumnCount]int) string {
	return headerContainerStyle.MarginLeft(lipgloss.Width(cursorView(false))).
		Render(lipgloss.JoinHorizontal(
			lipgloss.Center,
			renderRowColumn(
				headerStyle.Render("Package Name"),
				columnWidths[0],
				true,
			),
			renderRowColumn(headerStyle.Render("Wanted"), columnWidths[1], true),
			renderRowColumn(headerStyle.Render("Latest"), columnWidths[2], true),
			renderRowColumn(headerStyle.Render("Current"), columnWidths[3], false),
		))
}

func cursorView(active bool) string {
	c := ""
	if active {
		c = ">"
	}
	return cursorStyle.Width(2).Render(c)
}
