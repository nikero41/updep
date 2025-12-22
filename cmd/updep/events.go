package main

import (
	"errors"
	"fmt"
	"time"

	"updep/pkg/components/row"
	packagemanager "updep/pkg/packageManager"

	tea "github.com/charmbracelet/bubbletea"
)

type OutdatedPackagesCmd []packagemanager.Package

func getOutdatedPackages() tea.Msg {
	result, err := pm.GetOutdated()
	if err != nil {
		panic(err)
	}

	packages := []packagemanager.Package{}
	for packageName, value := range result {
		pkg, err := packagemanager.NewPackage(
			packageName,
			value.Wanted,
			value.Latest,
			value.Current,
		)
		if err != nil {
			_ = errors.New("invalid package versions")
			// TODO: handle error
			continue
		}

		packages = append(packages, *pkg)
	}

	return OutdatedPackagesCmd(packages)
}

type UpdateResultCmd struct {
	success bool
}

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
		return UpdateResultCmd{
			success: true,
		}
	}
}
