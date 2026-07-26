package update

import (
	"fmt"
	"strings"

	"updep/pkg/config"
	"updep/pkg/dependency"
	"updep/pkg/helpers"
	packagemanager "updep/pkg/packageManager"
	"updep/pkg/packageManager/events"
	"updep/pkg/version"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Update struct {
	spinner      spinner.Model
	Dependencies []dependency.Dependency
	Pm           packagemanager.PackageManager
	outputChan   chan events.PmOutputEvent
	output       []string
	isDone       bool
}

func New() *Update {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(config.Theme.Primary)

	return &Update{
		spinner:    s,
		outputChan: make(chan events.PmOutputEvent),
		isDone:     false,
	}
}

func (u Update) Init() tea.Cmd {
	u.Pm.Update(u.Dependencies, u.outputChan)
	return tea.Batch(u.spinner.Tick, helpers.WaitForChan(u.outputChan))
}

func (u *Update) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case helpers.ChannelMsg[events.PmOutputEvent]:
		u.output = append(u.output, msg.Result.Output)
		if msg.Result.Done {
			u.isDone = true
			close(u.outputChan)
			cmds = append(cmds, tea.Quit)
		} else {
			cmds = append(cmds, helpers.WaitForChan(u.outputChan))
		}
	}

	var cmd tea.Cmd
	u.spinner, cmd = u.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (u Update) Render() string {
	if !u.isDone {
		spinner := fmt.Sprintf("%v %s", u.spinner.View(), "Updating packages")
		return lipgloss.JoinVertical(lipgloss.Left, strings.Join(u.output, "\n"), spinner)
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

	report := fmt.Sprintf(
		"%v Updated %d packages: %v",
		checkMarkStyle.Render("✓"),
		len(u.Dependencies),
		strings.Join(depNames, ", "),
	)

	return lipgloss.JoinVertical(lipgloss.Left, strings.Join(u.output, "\n"), report)
}
