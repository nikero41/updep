package row

import (
	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Row struct {
	Pkg          packagemanager.Package
	Target       *packagemanager.Version
	ColumnWidths [config.ColumnCount]int
}

func New(pkg packagemanager.Package, columnWidths [config.ColumnCount]int) Row {
	return Row{
		Pkg:          pkg,
		ColumnWidths: columnWidths,
	}
}

func (r Row) Init() tea.Cmd {
	return nil
}

func (r Row) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, r.handleKeyPress(msg))
	}

	return r, tea.Batch(cmds...)
}

func (r Row) View() string {
	cells := make([]string, config.ColumnCount)
	for i, cell := range r.getCellStyles() {
		width := r.ColumnWidths[i] + config.ColumnGap
		if i == len(r.ColumnWidths)-1 {
			width -= config.ColumnGap
		}
		cells[i] = lipgloss.PlaceHorizontal(width, lipgloss.Left, cell)
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}

func (r Row) getCellStyles() [config.ColumnCount]string {
	var nameCellStyle, wantedCellStyle, latestCellStyle lipgloss.Style

	if r.Pkg.Current.Compare(r.Pkg.Wanted) == -1 {
		nameCellStyle = needUpdateStyle
	} else if r.Pkg.Current.Compare(r.Pkg.Wanted) == 1 {
		nameCellStyle = errorVersionStyle
	} else {
		nameCellStyle = optionalUpdateStyle
	}

	if r.Target != nil {
		switch *r.Target {
		case r.Pkg.Wanted:
			wantedCellStyle = wantedStyle
		case r.Pkg.Latest:
			latestCellStyle = latestStyle
		}
	}

	return [config.ColumnCount]string{
		nameCellStyle.Render(r.Pkg.Name),
		wantedCellStyle.Render(r.Pkg.Wanted.String()),
		latestCellStyle.Render(r.Pkg.Latest.String()),
		r.Pkg.Current.String(),
	}
}
