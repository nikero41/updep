package pkgtable

import (
	"updep/pkg/dependency"
	"updep/pkg/device"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up, Down, ExpandHelp, Homepage, MarkWanted, MarkLatest, ToggleTarget, Submit key.Binding
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

func (t *PkgTable) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Up):
		if t.cursor > 0 {
			t.cursor -= 1
		}

	case key.Matches(msg, keyMap.Down):
		if t.cursor < len(t.Packages)-1 {
			t.cursor += 1
		}

	case key.Matches(msg, keyMap.Homepage):
		pkg := t.Packages[t.cursor]
		err := device.OpenURL(pkg.Homepage)
		if err != nil {
			panic(err)
		}

	case key.Matches(msg, keyMap.MarkWanted):
		pkg := &t.Packages[t.cursor]
		if pkg.Current != pkg.Wanted {
			pkg.Target = dependency.Wanted
		}

	case key.Matches(msg, keyMap.MarkLatest):
		pkg := &t.Packages[t.cursor]
		if pkg.Current != pkg.Latest {
			pkg.Target = dependency.Latest
		}

	case key.Matches(msg, keyMap.ToggleTarget):
		pkg := &t.Packages[t.cursor]
		if pkg.Target != dependency.None {
			pkg.Target = dependency.None
			break
		}

		if pkg.Current.Compare(pkg.Wanted) >= 0 {
			pkg.Target = dependency.Latest
		} else {
			pkg.Target = dependency.Wanted
		}

	case key.Matches(msg, keyMap.Submit):
		var pkgs []dependency.Dependency
		for _, pkg := range t.Packages {
			if pkg.Target != dependency.None {
				pkgs = append(pkgs, pkg)
			}
		}
		return func() tea.Msg { return SelectPackagesMsg(pkgs) }

	case key.Matches(msg, keyMap.ExpandHelp):
		t.helpModel.ShowAll = !t.helpModel.ShowAll
	}

	return nil
}
