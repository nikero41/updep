package ui

import tea "charm.land/bubbletea/v2"

type Component interface {
	Update(msg tea.Msg) tea.Cmd
	Render() string
}
