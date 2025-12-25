package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit key.Binding
}

var keyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q/esc", "quit"),
	),
}

func (m *AppModel) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Quit):
		return tea.Quit
	}

	return nil
}
