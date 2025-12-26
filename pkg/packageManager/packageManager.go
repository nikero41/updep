package packagemanager

import (
	"errors"
	"os"

	"updep/pkg/packageManager/adapters"
	"updep/pkg/packages"
)

type PackageManager interface {
	GetOutdated() ([]packages.Package, error)
	Update(packages []packages.Package) error
}

func GetProjectPackageManager() (PackageManager, error) {
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

	if len(projectPms) == 1 {
		return projectPms[0], nil
	}

	if len(projectPms) == 0 {
		return nil, errors.New("no package manager found")
	}

	return nil, errors.New("multiple package managers found")
}
