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
		renderedRow := renderRow(pkg, columnWidths)
		if i == t.cursor {
			renderedRows[i] = activeRowStyle.Render("> " + renderedRow)
		} else {
			renderedRows[i] = "  " + renderedRow
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerView(columnWidths),
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
		t.helpModel.View(keyMap),
	)
}

func headerView(columnWidths [config.ColumnCount]int) string {
	columnNames := [config.ColumnCount]string{
		"Package Name",
		"Wanted",
		"Latest",
		"Current",
	}

	for i, name := range columnNames {
		gap := config.ColumnGap
		if i == len(columnNames)-1 {
			gap = 0
		}
		columnNames[i] = lipgloss.PlaceHorizontal(
			columnWidths[i]+gap,
			lipgloss.Left,
			name,
		)
	}

	return "  " + headerStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		columnNames[:]...,
	))
}
