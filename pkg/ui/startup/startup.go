package startup

import (
	"fmt"

	"updep/pkg/config"
	packagemanager "updep/pkg/packageManager"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// TODO: handle not installed dependencies

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

func (s *StartUp) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.pmListModel.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		cmds = append(cmds, s.handleKeyPress(msg))

	case PackageManagersFoundCmd:
		switch len(msg) {
		case 0:
			cmds = append(cmds, tea.Sequence(tea.Println("All packages are up-to-date"), tea.Quit))
		case 1:
			s.pm = msg[0]
			cmds = append(cmds, getOutdatedDependencies(s.pm))
		default:
			cmds = append(cmds, s.pmListModel.SetItems(msg))
		}

	case OutdatedDependenciesMsg:
		return func() tea.Msg {
			return StartUpCompletedMsg{
				Dependencies: msg,
				Pm:           s.pm,
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

	return tea.Batch(cmds...)
}

func (s StartUp) Render() string {
	if s.pmListModel.Items() != nil {
		return s.pmListModel.View()
	}
	return fmt.Sprintf("%v %s", s.spinner.View(), s.labelText)
}
