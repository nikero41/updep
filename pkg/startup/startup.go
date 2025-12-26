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

// New creates a StartUp model with a configured spinner and default label.
// The spinner uses the Points glyph set and is styled with the configured theme's Mauve color; the returned StartUp's labelText is "Getting outdated packages" and pm is nil.
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
	case OutdatedPackagesMsg:
		return s, func() tea.Msg {
			return StartUpCompletedMsg{
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