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

// New creates a PkgTable with an initialized help model, the cursor set to 0, and an empty Packages slice.
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
		if i == t.cursor {
			renderedRows[i] = activeRowStyle.Render("> " + row)
		} else {
			renderedRows[i] = "  " + row
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerView(columnWidths),
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
		t.helpModel.View(keyMap),
	)
}

// headerView renders the table header row using the provided column widths.
// 
// columnWidths is an array of widths for the columns in order: Package Name, Wanted, Latest, Current.
// The result is a styled header string, positioned with a left margin and spacing appropriate for the given widths.
func headerView(columnWidths [config.ColumnCount]int) string {
	return headerStyle.MarginLeft(2).Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.PlaceHorizontal(
			columnWidths[0]+config.ColumnGap,
			lipgloss.Left,
			"Package Name",
		),
		lipgloss.PlaceHorizontal(
			columnWidths[1]+config.ColumnGap,
			lipgloss.Left,
			"Wanted",
		),
		lipgloss.PlaceHorizontal(
			columnWidths[2]+config.ColumnGap,
			lipgloss.Left,
			"Latest",
		),
		lipgloss.PlaceHorizontal(
			columnWidths[3],
			lipgloss.Left,
			"Current",
		),
	))
}