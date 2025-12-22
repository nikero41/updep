package row

import (
	"fmt"

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
	var nameCellStyle lipgloss.Style

	if r.Target != nil {
		switch *r.Target {
		case r.Pkg.Wanted:
			nameCellStyle = warningStyle
		case r.Pkg.Latest:
			nameCellStyle = errorStyle
		}
	}

	wantedVersionDiff := r.Pkg.Current.Diff(r.Pkg.Wanted)
	latestVersionDiff := r.Pkg.Current.Diff(r.Pkg.Latest)

	return [config.ColumnCount]string{
		nameCellStyle.Render(r.Pkg.Name),
		versionWithDiff(r.Pkg.Wanted, wantedVersionDiff),
		versionWithDiff(r.Pkg.Latest, latestVersionDiff),
		r.Pkg.Current.String(),
	}
}

func versionWithDiff(
	v packagemanager.Version,
	diff packagemanager.Version,
) string {
	if diff.Major > 0 {
		return errorStyle.Render(v.String())
	} else if diff.Minor > 0 {
		return fmt.Sprintf("%d.%v", v.Major, warningStyle.Render(fmt.Sprintf("%d.%d", v.Minor, v.Patch)))
	} else if diff.Patch > 0 {
		return fmt.Sprintf("%d.%d.%v", v.Major, v.Minor, successStyle.Render(fmt.Sprintf("%d", v.Patch)))
	}

	return v.String()
}
