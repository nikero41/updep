package main

import (
	"fmt"
	"time"

	"updep/pkg/components/row"
	"updep/pkg/packages"

	tea "github.com/charmbracelet/bubbletea"
)

type OutdatedPackagesCmd []packages.Package

func getOutdatedPackages() tea.Msg {
	packages, err := pm.GetOutdated()
	if err != nil {
		panic(err)
	}

	return OutdatedPackagesCmd(packages)
}

type UpdateResultCmd bool

func updatePackages(_ []row.Row) tea.Cmd {
	return func() tea.Msg {
		timer := time.NewTimer(time.Second * 5)
		<-timer.C
		fmt.Println("🪚 timer.C:", timer.C)
		// output, err := exec.Command("npm", "update").Output()
		// if err != nil {
		// 	fmt.Println("🪚 err:", err)
		// }
		// fmt.Println("🪚 output:", output)

		// TODO: investigate why returning tea.Quit from update Msg is not quiting
		return UpdateResultCmd(true)
	}
}
