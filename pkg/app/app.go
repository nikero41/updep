package app

import (
	"updep/pkg/startup"

	tea "github.com/charmbracelet/bubbletea"

	packagemanager "updep/pkg/packageManager"

	"github.com/charmbracelet/bubbles/help"
)

type AppScreen int

const (
	StartUp AppScreen = iota
	Choice
)

type AppModel struct {
	startUpModel  startup.StartUp
	screen        AppScreen
	pm            packagemanager.PackageManager
}

func New() AppModel {
	return AppModel{
		startUpModel:  startup.New(),
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
	case startup.StartUpCompletedCmd:
		m.screen = Choice

	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))

	case UpdateResultCmd:
		return m, tea.Quit
	}

	switch m.screen {
	case StartUp:
		newModel, cmd := m.startUpModel.Update(msg)
		m.startUpModel = newModel.(startup.StartUp)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	switch m.screen {
	case StartUp:
		return m.startUpModel.View()
	default:
		return ""
	}
}
