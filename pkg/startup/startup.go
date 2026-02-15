package startup

import (
	"fmt"

	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"
	"updep/pkg/packageManager/adapters"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartUp struct {
	spinner     spinner.Model
	labelText   string
	pm          packagemanager.PackageManager
	pmListModel PmList
}

func New() *StartUp {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Primary)

	return &StartUp{
		spinner:     s,
		labelText:   "Getting outdated packages",
		pm:          nil,
		pmListModel: NewPmList("Choose a package manager"),
	}
}

func (s StartUp) Init() tea.Cmd {
	return tea.Batch(s.spinner.Tick, getPackageManager())
}

func (s StartUp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.pmListModel.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		cmds = append(cmds, s.handleKeyPress(msg))

	case PackageManagersFoundCmd:
		// TODO: temp for dev
		if len(msg) == 0 {
			msg = append(msg, adapters.NewNpm())
		}

		switch len(msg) {
		case 0:
			// TODO: Show error screen
			panic("not implemented")
		case 1:
			s.pm = msg[0]
			cmds = append(cmds, getOutdatedDependencies(s.pm))
		default:
			cmds = append(cmds, s.pmListModel.SetItems(msg))
		}

	case OutdatedDependenciesMsg:
		return s, func() tea.Msg {
			return StartUpCompletedMsg{
				Dependencies: msg,
				Pm:       s.pm,
			}
		}
	}

	var cmd tea.Cmd
	if s.pmListModel.Items() != nil {
		s.pmListModel.Model, cmd = s.pmListModel.Update(msg)
	} else {
		s.spinner, cmd = s.spinner.Update(msg)
	}
	cmds = append(cmds, cmd)

	return s, tea.Batch(cmds...)
}

func (s StartUp) View() string {
	if s.pmListModel.Items() != nil {
		return s.pmListModel.View()
	}

	return fmt.Sprintf("%v %s", s.spinner.View(), s.labelText)
}
