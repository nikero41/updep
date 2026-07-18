package app

import (
	depstable "updep/pkg/ui/depsTable"
	"updep/pkg/ui/startup"
	"updep/pkg/ui/update"

	tea "charm.land/bubbletea/v2"

	packagemanager "updep/pkg/packageManager"
)

type AppScreen int

const (
	StartUp AppScreen = iota
	Choice
	Update
)

type AppModel struct {
	startUpModel   *startup.StartUp
	depsTableModel *depstable.DepsTable
	updateModel    *update.Update
	screen         AppScreen
	pm             packagemanager.PackageManager
	height         int
}

func New(pm packagemanager.PackageManager) *AppModel {
	return &AppModel{
		startUpModel:   startup.New(),
		depsTableModel: depstable.New(),
		updateModel:    update.New(),
		screen:         StartUp,
		pm:             pm,
		height:         0,
	}
}

func (m AppModel) Init() tea.Cmd {
	return m.startUpModel.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.depsTableModel.SetHeight(msg.Height)

	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))

	case startup.StartUpCompletedMsg:
		m.screen = Choice
		m.depsTableModel.SetDependencies(msg.Dependencies)
		m.pm = msg.Pm
		cmds = append(cmds, m.depsTableModel.Init())

	case depstable.SelectDependenciesMsg:
		m.screen = Update
		m.updateModel.Pm = m.pm
		m.updateModel.Dependencies = msg
		cmds = append(cmds, m.updateModel.Init())
	}

	switch m.screen {
	case StartUp:
		cmd := m.startUpModel.Update(msg)
		cmds = append(cmds, cmd)
	case Choice:
		cmd := m.depsTableModel.Update(msg)
		cmds = append(cmds, cmd)
	case Update:
		cmd := m.updateModel.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() tea.View {
	var content string
	switch m.screen {
	case StartUp:
		content = m.startUpModel.Render()
	case Choice:
		content = m.depsTableModel.Render()
	case Update:
		content = m.updateModel.Render()
	default:
		content = ""
	}

	return tea.NewView(content)
}
