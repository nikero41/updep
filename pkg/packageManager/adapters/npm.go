package adapters

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"

	"updep/pkg/packages"
)

type Npm struct {
	name string
}

func NewNpm() Npm {
	return Npm{name: "npm"}
}

type JSONPackage struct {
	Name    string
	Wanted  string `json:"wanted"`
	Latest  string `json:"latest"`
	Current string `json:"current"`
}

func (pm Npm) Name() string { return pm.name }

func (pm Npm) GetOutdated() ([]packages.Package, error) {
	output, err := exec.Command("npm", "outdated", "--json", "--long").Output()

	// NOTE: npm outdated returns exit status 1 if there are outdated packages
	if err != nil && err.Error() != "exit status 1" {
		return nil, err
	}

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	outdatedPackages := make([]packages.Package, len(outdated))

	var index int
	for name, value := range outdated {
		pkg, err := packages.New(
			name,
			value.Wanted,
			value.Latest,
			value.Current,
		)
		if err != nil {
			return nil, errors.New("invalid package versions")
		}

		outdatedPackages[index] = *pkg
		index++
	}

	return outdatedPackages, nil
}

func (pm Npm) Update(packages []packages.Package) error {
	return errors.New("not implemented")
}
