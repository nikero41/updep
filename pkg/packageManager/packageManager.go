package packagemanager

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"updep/pkg/packageManager/adapters"
	"updep/pkg/dependency"
)

type PackageManager interface {
	Name() string
	GetOutdated() ([]dependency.Dependency, error)
	Update(packages []dependency.Dependency) ([]byte, error)
}

var lockfilePmMapper = map[string]string{
	"package-lock.json": "npm",
}

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
			if pmName := lockfilePmMapper[file.Name()]; pmName != "" {
				pm, err := New(pmName)
				if err != nil {
					return nil, err
				}
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
	pm, err := New(pmName)
	if err != nil {
		return nil, err
	}

	return pm, nil
}

func New(pmName string) (PackageManager, error) {
	switch pmName {
	case "npm":
		return adapters.NewNpm(), nil
	case "":
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid package manager: %s", pmName)
	}
}
