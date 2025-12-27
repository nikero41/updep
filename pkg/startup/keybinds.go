package startup

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Select key.Binding
}

var keyMap = KeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("ent", "Select"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Select}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Select},
	}
}

func (s *StartUp) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Select):
		if i, ok := s.pmListModel.SelectedItem().(PmListItem); ok {
			s.pm = i.pm
			return getOutdatedPackages(s.pm)
		}
	}

	return nil
}
