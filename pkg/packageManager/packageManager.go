package packagemanager

import (
	"fmt"
	"os"

	"updep/pkg/packageManager/adapters"
	"updep/pkg/packages"
)

type PackageManager interface {
	GetOutdated() ([]packages.Package, error)
	Update(packages []packages.Package) error
}

func GetProjectPackageManager() PackageManager {
	dir, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var projectPms []PackageManager
	for _, file := range dir {
		switch file.Name() {
		case "package-lock.json":
			projectPms = append(projectPms, adapters.Npm{})
		}
	}
	fmt.Println("🪚 pmLockFiles:", projectPms)

	return adapters.Npm{}
}
