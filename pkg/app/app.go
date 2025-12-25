package app

import (
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
	screen        AppScreen
}

func New() AppModel {
	return AppModel{
		screen:        StartUp,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))

	case UpdateResultCmd:
		return m, tea.Quit
	}

	switch m.screen {
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	switch m.screen {
	default:
		return ""
	}
}
