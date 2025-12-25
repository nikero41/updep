package main

import (
	"updep/pkg/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(app.New())
	defer p.Kill()
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
