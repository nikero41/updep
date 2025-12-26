package main

import (
	"updep/pkg/app"

	tea "github.com/charmbracelet/bubbletea"
)

// main starts the Bubble Tea TUI returned by app.New.
// It defers program cleanup via Kill and panics if running the program fails.
func main() {
	p := tea.NewProgram(app.New())
	defer p.Kill()
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}