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

// getPackageManager returns a tea.Cmd that retrieves the project's package manager and emits a PackageManagerFoundCmd.
// The command panics if retrieving the package manager fails.
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

// getOutdatedPackages returns a tea.Cmd that retrieves outdated packages from pm and sends them as an OutdatedPackagesMsg.
// The command panics if pm.GetOutdated returns an error.
func getOutdatedPackages(pm packagemanager.PackageManager) tea.Cmd {
	return func() tea.Msg {
		packages, err := pm.GetOutdated()
		if err != nil {
			panic(err)
		}

		return OutdatedPackagesMsg(packages)
	}
}