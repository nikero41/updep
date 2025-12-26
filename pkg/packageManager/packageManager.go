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

// GetProjectPackageManager detects the project's package manager by inspecting the current directory.
// It returns the detected PackageManager when exactly one manager is found.
// If no package manager files are present, it returns an error "no package manager found".
// If multiple manager indicators are present, it returns an error "multiple package managers found".
// Detection currently maps the presence of "package-lock.json" to adapters.Npm{} and any filesystem read error is returned directly.
func GetProjectPackageManager() (PackageManager, error) {
	dir, err := os.ReadDir(".")
	if err != nil {
		return nil, err
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