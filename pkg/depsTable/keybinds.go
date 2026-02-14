package pkgtable

import (
	"updep/pkg/dependency"
	"updep/pkg/device"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up, Down, ExpandHelp, Homepage,
	MarkWanted, MarkLatest, ToggleTarget, SelectAll, InvertSelection,
	Submit key.Binding
}

var keyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	ExpandHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),

	Homepage: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "view in browser"),
	),

	// Version selection
	MarkWanted: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "update to wanted version"),
	),
	MarkLatest: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "update to latest version"),
	),
	ToggleTarget: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle"),
	),
	SelectAll: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "select all"),
	),
	InvertSelection: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "invert selection"),
	),

	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("ent", "Update selected"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.ExpandHelp}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Up,
			// TODO: add help from app
			// k.Quit,
			k.ToggleTarget,
			k.Homepage,
		},
		{k.Down, k.MarkWanted},
		{k.ExpandHelp, k.MarkLatest},
	}
}

func (t *DepsTable) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Up):
		if t.cursor > 0 {
			t.cursor -= 1
		}

	case key.Matches(msg, keyMap.Down):
		if t.cursor < len(t.Dependencies)-1 {
			t.cursor += 1
		}

	case key.Matches(msg, keyMap.Homepage):
		err := device.OpenURL(t.Dependencies[t.cursor].Homepage)
		if err != nil {
			panic(err)
		}

	case key.Matches(msg, keyMap.MarkWanted):
		d := &t.Dependencies[t.cursor]
		if d.Current != d.Wanted {
			d.Target = dependency.Wanted
		}

	case key.Matches(msg, keyMap.MarkLatest):
		d := &t.Dependencies[t.cursor]
		if d.Current != d.Latest {
			d.Target = dependency.Latest
		}

	case key.Matches(msg, keyMap.ToggleTarget):
		t.Dependencies[t.cursor].ToggleTarget()

	case key.Matches(msg, keyMap.SelectAll):
		shouldClear := true
		for i := range t.Dependencies {
			if t.Dependencies[i].Target == dependency.None {
				shouldClear = false
				t.Dependencies[i].SelectTarget()
			}
		}

		if shouldClear {
			for i := range t.Dependencies {
				t.Dependencies[i].Target = dependency.None
			}
		}

	case key.Matches(msg, keyMap.InvertSelection):
		for i := range t.Dependencies {
			t.Dependencies[i].ToggleTarget()
		}

	case key.Matches(msg, keyMap.Submit):
		var deps []dependency.Dependency
		for _, d := range t.Dependencies {
			if d.Target != dependency.None {
				deps = append(deps, d)
			}
		}
		return func() tea.Msg { return SelectDependenciesMsg(deps) }

	case key.Matches(msg, keyMap.ExpandHelp):
		t.helpModel.ShowAll = !t.helpModel.ShowAll
	}

	return nil
}
