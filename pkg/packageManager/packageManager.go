package packagemanager

import (
	"updep/pkg/packageManager/adapters"
	"updep/pkg/packages"
)

type PackageManager interface {
	GetOutdated() ([]packages.Package, error)
	Update(packages []packages.Package) error
}

func GetProjectPackageManager() PackageManager {
	return adapters.Npm{}
}
