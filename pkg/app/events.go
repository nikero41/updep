package app

import (
	"fmt"
	"time"

	"updep/pkg/packages"

	tea "github.com/charmbracelet/bubbletea"
)

type UpdateResultCmd bool

func updatePackages(_ []packages.Package) tea.Cmd {
	return func() tea.Msg {
		timer := time.NewTimer(time.Second * 5)
		<-timer.C
		fmt.Println("🪚 timer.C:", timer.C)
		// output, err := exec.Command("npm", "update").Output()
		// if err != nil {
		// 	panic(err)
		// }

		// TODO: investigate why returning tea.Quit from update Msg is not quiting
		return UpdateResultCmd(true)
	}
}
