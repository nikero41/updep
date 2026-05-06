package startup

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
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
			return getOutdatedDependencies(s.pm)
		}
	}

	return nil
}
