package app

import (
	pkgtable "updep/pkg/pkgTable"
	"updep/pkg/startup"

	tea "github.com/charmbracelet/bubbletea"

	packagemanager "updep/pkg/packageManager"
)

type AppScreen int

const (
	StartUp AppScreen = iota
	Choice
)

type AppModel struct {
	startUpModel  startup.StartUp
	pkgTableModel pkgtable.PkgTable
	screen        AppScreen
	pm            packagemanager.PackageManager
}

func New() AppModel {
	return AppModel{
		startUpModel:  startup.New(),
		pkgTableModel: pkgtable.New(),
		screen:        StartUp,
		pm:            nil,
	}
}

func (m AppModel) Init() tea.Cmd {
	return m.startUpModel.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case startup.StartUpCompletedMsg:
		m.screen = Choice
		m.pkgTableModel.Packages = msg.Packages
		cmd := m.pkgTableModel.Init()
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))
	}

	switch m.screen {
	case StartUp:
		newModel, cmd := m.startUpModel.Update(msg)
		m.startUpModel = newModel.(startup.StartUp)
		cmds = append(cmds, cmd)
	case Choice:
		newModel, cmd := m.pkgTableModel.Update(msg)
		m.pkgTableModel = newModel.(pkgtable.PkgTable)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	switch m.screen {
	case StartUp:
		return m.startUpModel.View()
	case Choice:
		return m.pkgTableModel.View()
	default:
		return ""
	}
}
