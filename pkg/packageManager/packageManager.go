package packagemanager

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"

	"updep/pkg/packageManager/adapters"
	"updep/pkg/packages"
)

type PackageManager interface {
	Name() string
	GetOutdated() ([]packages.Package, error)
	Update(packages []packages.Package) error
}

var (
	lockfilePmMapper = map[string]PackageManager{
		"package-lock.json": adapters.NewNpm(),
	}
	namePmMapper = map[string]PackageManager{
		"npm": adapters.NewNpm(),
	}
)

func GetProjectPackageManagers() ([]PackageManager, error) {
	dir, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var projectPms []PackageManager
	for _, file := range dir {
		switch file.Name() {
		case "package.json":
			pm, err := getPackageJSONPackageManager()
			if err != nil {
				return nil, err
			}
			if pm != nil {
				slog.Info("package manager in package.json", "packageManager", pm)
				return []PackageManager{pm}, nil
			}

		default:
			if pm := lockfilePmMapper[file.Name()]; pm != nil {
				projectPms = append(projectPms, pm)
			}
		}
	}

	slog.Info("package managers found", "packageManagers", projectPms)
	return projectPms, nil
}

func getPackageJSONPackageManager() (PackageManager, error) {
	var packageJSON struct {
		PackageManager string `json:"packageManager"`
	}

	packageJSONFile, err := os.ReadFile("package.json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(strings.NewReader(string(packageJSONFile))).
		Decode(&packageJSON)
	if err != nil {
		return nil, err
	}

	pmName := strings.Split(packageJSON.PackageManager, "@")[0]

	if pm := namePmMapper[pmName]; pm != nil {
		return pm, nil
	}

	return nil, nil
}
