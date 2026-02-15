package app

import (
	depstable "updep/pkg/depsTable"
	"updep/pkg/startup"
	"updep/pkg/update"

	tea "github.com/charmbracelet/bubbletea"

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

func New() *AppModel {
	return &AppModel{
		startUpModel:   startup.New(),
		depsTableModel: depstable.New(),
		updateModel:    update.New(),
		screen:         StartUp,
		pm:             nil,
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
		newModel, cmd := m.startUpModel.Update(msg)
		newStartUpModel := newModel.(startup.StartUp)
		m.startUpModel = &newStartUpModel
		cmds = append(cmds, cmd)
	case Choice:
		newModel, cmd := m.depsTableModel.Update(msg)
		newDepsTableModel := newModel.(depstable.DepsTable)
		m.depsTableModel = &newDepsTableModel
		cmds = append(cmds, cmd)
	case Update:
		newModel, cmd := m.updateModel.Update(msg)
		newUpdateModel := newModel.(update.Update)
		m.updateModel = &newUpdateModel
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	switch m.screen {
	case StartUp:
		return m.startUpModel.View()
	case Choice:
		return m.depsTableModel.View()
	case Update:
		return m.updateModel.View()
	default:
		return ""
	}
}
