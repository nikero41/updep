package update

import (
	"log"

	"updep/pkg/dependency"
	packagemanager "updep/pkg/packageManager"

	tea "charm.land/bubbletea/v2"
)

type UpdateCompleteMsg struct {
	output  []byte
	isError bool
}

func updateDependencies(
	pm packagemanager.PackageManager,
	deps []dependency.Dependency,
) tea.Cmd {
	return func() tea.Msg {
		// TODO: create backup of package.json and package-lock.json

		output, err := pm.Update(deps)
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
