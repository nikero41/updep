package update

import (
	"log"

	packagemanager "updep/pkg/packageManager"
	"updep/pkg/packages"

	tea "github.com/charmbracelet/bubbletea"
)

type UpdateCompleteMsg struct {
	output  []byte
	isError bool
}

func updatePackages(
	pm packagemanager.PackageManager,
	packages []packages.Package,
) tea.Cmd {
	return func() tea.Msg {
		// TODO: create backup of package.json and package-lock.json

		output, err := pm.Update(packages)
		log.Println("🪚 output:", string(output))
		if err != nil {
			log.Println("🪚 err:", err)
			// TODO: restore from backup of package.json and package-lock.json
			panic(err)
		}

		// TODO: remove backup of package.json and package-lock.json

		return UpdateCompleteMsg{
			output:  output,
			isError: false,
		}
	}
}
