package pkgtable

import (
	"updep/pkg/config"
	"updep/pkg/packages"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PkgTable struct {
	helpModel help.Model
	cursor    int
	Packages  []packages.Package
}

func New() PkgTable {
	return PkgTable{
		helpModel: help.New(),
		cursor:    0,
		Packages:  []packages.Package{},
	}
}

func (t PkgTable) Init() tea.Cmd {
	return nil
}

func (t PkgTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, t.handleKeyPress(msg))
	}

	return t, tea.Batch(cmds...)
}

func (t PkgTable) View() string {
	columnWidths := calculateColumnWidths(t.Packages)

	renderedRows := make([]string, len(t.Packages))
	for i, pkg := range t.Packages {
		row := renderRow(pkg, columnWidths)
		renderedRows[i] = cursorView(i == t.cursor) + row
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerView(columnWidths),
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
		helpContainerStyle.Render(t.helpModel.View(keyMap)),
	)
}

func headerView(columnWidths [config.ColumnCount]int) string {
	return headerContainerStyle.MarginLeft(lipgloss.Width(cursorView(false))).
		Render(lipgloss.JoinHorizontal(
			lipgloss.Center,
			renderHeaderColumn("Package Name", columnWidths[0], true),
			renderHeaderColumn("Wanted", columnWidths[1], true),
			renderHeaderColumn("Latest", columnWidths[2], true),
			renderHeaderColumn("Current", columnWidths[3], false),
		))
}

func renderHeaderColumn(label string, width int, withGap bool) string {
	if withGap {
		width += config.ColumnGap
	}

	return lipgloss.PlaceHorizontal(
		width,
		lipgloss.Left,
		headerStyle.Render(label),
	)
}

func cursorView(active bool) string {
	c := ""
	if active {
		c = ">"
	}
	return cursorStyle.Width(2).Render(c)
}
