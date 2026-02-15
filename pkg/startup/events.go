package startup

import (
	"slices"
	"strings"

	"updep/pkg/dependency"
	packagemanager "updep/pkg/packageManager"

	tea "github.com/charmbracelet/bubbletea"
)

type StartUpCompletedMsg struct {
	Dependencies []dependency.Dependency
	Pm           packagemanager.PackageManager
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

type OutdatedDependenciesMsg []dependency.Dependency

func getOutdatedDependencies(pm packagemanager.PackageManager) tea.Cmd {
	return func() tea.Msg {
		outdatedDeps, err := pm.GetOutdated()
		if err != nil {
			panic(err)
		}

		slices.SortStableFunc(outdatedDeps, func(a, b dependency.Dependency) int {
			return strings.Compare(a.Name, b.Name)
		})

		return OutdatedDependenciesMsg(outdatedDeps)
	}
}
