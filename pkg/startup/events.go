package startup

import (
	packagemanager "updep/pkg/packageManager"
	"updep/pkg/packages"

	tea "github.com/charmbracelet/bubbletea"
)

type StartUpCompletedMsg struct {
	Packages []packages.Package
	Pm       packagemanager.PackageManager
}

type PackageManagerFoundCmd packagemanager.PackageManager

func getPackageManager() tea.Cmd {
	return func() tea.Msg {
		pm, err := packagemanager.GetProjectPackageManager()
		if err != nil {
			panic(err)
		}

		return PackageManagerFoundCmd(pm)
	}
}

type OutdatedPackagesMsg []packages.Package

func getOutdatedPackages(pm packagemanager.PackageManager) tea.Cmd {
	return func() tea.Msg {
		packages, err := pm.GetOutdated()
		if err != nil {
			panic(err)
		}

		return OutdatedPackagesMsg(packages)
	}
}
