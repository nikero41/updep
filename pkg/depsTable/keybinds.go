package depstable

import (
	"updep/pkg/dependency"
	"updep/pkg/device"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up, Down, ToTop, ToBottom, HalfPageUp, HalfPageDown,
	ExpandHelp, Homepage,
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
	ToTop: key.NewBinding(
		key.WithKeys("g", tea.KeyHome.String()), // TODO: make it gg
		key.WithHelp("g", "top"),
	),
	ToBottom: key.NewBinding(
		key.WithKeys("G", tea.KeyEnd.String()),
		key.WithHelp("G", "bottom"),
	),
	HalfPageUp: key.NewBinding(
		key.WithKeys(tea.KeyCtrlU.String()),
		key.WithHelp("^u", "scroll up"),
	),
	HalfPageDown: key.NewBinding(
		key.WithKeys(tea.KeyCtrlD.String()),
		key.WithHelp("^d", "scroll down"),
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
		t.scrollFix()

	case key.Matches(msg, keyMap.Down):
		if t.cursor < len(t.dependencies)-1 {
			t.cursor += 1
		}
		t.scrollFix()

	case key.Matches(msg, keyMap.ToTop):
		t.cursor = 0
		t.scrollFix()

	case key.Matches(msg, keyMap.ToBottom):
		t.cursor = len(t.dependencies) - 1
		t.scrollFix()

	case key.Matches(msg, keyMap.HalfPageUp):
		step := t.rowCount() / 2
		t.scrollTo(t.offset - step)

	case key.Matches(msg, keyMap.HalfPageDown):
		step := t.rowCount() / 2
		t.scrollTo(t.offset + step)

	case key.Matches(msg, keyMap.Homepage):
		err := device.OpenURL(t.dependencies[t.cursor].Homepage)
		if err != nil {
			panic(err)
		}

	case key.Matches(msg, keyMap.MarkWanted):
		d := &t.dependencies[t.cursor]
		if d.Current != d.Wanted {
			d.Target = dependency.Wanted
		}

	case key.Matches(msg, keyMap.MarkLatest):
		d := &t.dependencies[t.cursor]
		if d.Current != d.Latest {
			d.Target = dependency.Latest
		}

	case key.Matches(msg, keyMap.ToggleTarget):
		t.dependencies[t.cursor].ToggleTarget()

	case key.Matches(msg, keyMap.SelectAll):
		shouldClear := true
		for i := range t.dependencies {
			if t.dependencies[i].Target == dependency.None {
				shouldClear = false
				t.dependencies[i].SelectTarget()
			}
		}

		if shouldClear {
			for i := range t.dependencies {
				t.dependencies[i].Target = dependency.None
			}
		}

	case key.Matches(msg, keyMap.InvertSelection):
		for i := range t.dependencies {
			t.dependencies[i].ToggleTarget()
		}

	case key.Matches(msg, keyMap.Submit):
		var deps []dependency.Dependency
		for _, d := range t.dependencies {
			if d.Target != dependency.None {
				deps = append(deps, d)
			}
		}
		return func() tea.Msg { return SelectDependenciesMsg(deps) }

	case key.Matches(msg, keyMap.ExpandHelp):
		t.helpModel.ShowAll = !t.helpModel.ShowAll
		t.scrollFix()
	}

	return nil
}
