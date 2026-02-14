package startup

import (
	"slices"

	packagemanager "updep/pkg/packageManager"
	"updep/pkg/dependency"

	tea "github.com/charmbracelet/bubbletea"
)

type StartUpCompletedMsg struct {
	Packages []dependency.Dependency
	Pm       packagemanager.PackageManager
}

type PackageManagersFoundCmd []packagemanager.PackageManager

func getPackageManager() tea.Cmd {
	return func() tea.Msg {
		pms, err := packagemanager.GetProjectPackageManagers()
		if err != nil {
			panic(err)
		}

		return PackageManagersFoundCmd(pms)
	}
}

type OutdatedPackagesMsg []dependency.Dependency

func getOutdatedPackages(pm packagemanager.PackageManager) tea.Cmd {
	return func() tea.Msg {
		outdatedPackages, err := pm.GetOutdated()
		if err != nil {
			panic(err)
		}

		slices.SortStableFunc(outdatedPackages, func(a, b dependency.Dependency) int {
			if a.Name < b.Name {
				return -1
			}
			if a.Name > b.Name {
				return 1
			}
			return 0
		})

		return OutdatedPackagesMsg(outdatedPackages)
	}
}
