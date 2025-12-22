package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"updep/pkg/packages"
)

type Npm struct{}

func NewNpm() Npm {
	return Npm{}
}

type JSONPackage struct {
	Name    string
	Wanted  string `json:"wanted"`
	Latest  string `json:"latest"`
	Current string `json:"current"`
}

func (pm Npm) GetOutdated() ([]packages.Package, error) {
	output, err := exec.Command("npm", "outdated", "--json").Output()
	// output, err := os.ReadFile("output.json")
	if err != nil {
		fmt.Println("🪚 err:", err)
	}

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	outdatedPackages := make([]packages.Package, len(outdated))
	for name, value := range outdated {
		pkg, err := packages.NewPackage(
			name,
			value.Wanted,
			value.Latest,
			value.Current,
		)
		if err != nil {
			_ = errors.New("invalid package versions")
			// TODO: handle error
			continue
		}

		outdatedPackages = append(outdatedPackages, *pkg)
	}

	return outdatedPackages, nil
}

func (pm Npm) Update(packages []packages.Package) error {
	return errors.New("not implemented")
}
