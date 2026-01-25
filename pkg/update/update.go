package update

import (
	"fmt"

	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"
	"updep/pkg/packages"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Update struct {
	spinner   spinner.Model
	labelText string
	Packages  []packages.Package
	Pm        packagemanager.PackageManager
	output    []byte
}

func New() Update {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Primary)

	return Update{
		spinner:   s,
		labelText: "Getting outdated packages",
	}
}

func (u Update) Init() tea.Cmd {
	return tea.Batch(u.spinner.Tick, updatePackages(u.Pm, u.Packages))
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
	if u.output != nil {
		return fmt.Sprintf(
			"%s\n\n✓ Updated 3 packages: lodash, react, typescript",
			u.output,
		)
	}

	return fmt.Sprintf("%v %s", u.spinner.View(), u.labelText)
}
