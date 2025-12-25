package startup

import (
	"fmt"

	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartUp struct {
	spinner   spinner.Model
	labelText string
	pm        packagemanager.PackageManager
}

func New() StartUp {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Mauve)

	return StartUp{
		spinner:   s,
		labelText: "Getting outdated packages",
		pm:        nil,
	}
}

func (s StartUp) Init() tea.Cmd {
	return tea.Batch(s.spinner.Tick, getPackageManager())
}

func (s StartUp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case PackageManagerFoundCmd:
		s.pm = msg
		return s, getOutdatedPackages(msg)
	case OutdatedPackagesCmd:
		return s, func() tea.Msg {
			return StartUpCompletedCmd{
				Packages: msg,
				Pm:       s.pm,
			}
		}
	}

	var cmd tea.Cmd
	s.spinner, cmd = s.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return s, tea.Batch(cmds...)
}

func (s StartUp) View() string {
	return fmt.Sprintf("%v %s", s.spinner.View(), s.labelText)
}
