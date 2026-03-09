package update

import (
	"fmt"
	"strings"

	"updep/pkg/config"
	"updep/pkg/dependency"
	packagemanager "updep/pkg/packageManager"
	"updep/pkg/version"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Update struct {
	spinner      spinner.Model
	labelText    string
	Dependencies []dependency.Dependency
	Pm           packagemanager.PackageManager
	output       []byte
	isDone       bool
}

func New() *Update {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Primary)

	return &Update{
		spinner:   s,
		labelText: "Updating packages",
		isDone:    false,
	}
}

func (u Update) Init() tea.Cmd {
	return tea.Batch(u.spinner.Tick, updateDependencies(u.Pm, u.Dependencies))
}

func (u Update) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case UpdateCompleteMsg:
		u.output = msg.output
		cmds = append(cmds, tea.Quit)
	}

	var cmd tea.Cmd
	u.spinner, cmd = u.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return u, tea.Batch(cmds...)
}

func (u Update) View() string {
	if !u.isDone {
		return fmt.Sprintf("%v %s", u.spinner.View(), u.labelText)
	}

	var depNames []string
	for _, d := range u.Dependencies {
		var depStyle lipgloss.Style

		switch d.DiffLevel() {
		case version.Major:
			depStyle = majorDiffStyle
		case version.Minor:
			depStyle = minorDiffStyle
		case version.Patch:
			depStyle = patchDiffStyle
		}
		depNames = append(depNames, depStyle.Render(d.Name))
	}

	return fmt.Sprintf(
		"\n\n%v Updated %d packages: %v",
		checkMarkStyle.Render("✓"),
		len(u.Dependencies),
		strings.Join(depNames, ", "),
	)
}
