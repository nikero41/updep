package adapters

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"updep/pkg/dependency"

	"github.com/go-playground/validator/v10"
)

type Npm struct {
	name string
}

func NewNpm() *Npm {
	return &Npm{name: "npm"}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

type JSONPackage struct {
	Wanted   string `validate:"required,semver" json:"wanted"`
	Latest   string `validate:"required,semver" json:"latest"`
	Current  string `validate:"required,semver" json:"current"`
	Homepage string `validate:"omitempty,url"   json:"homepage"`
}

func (pm Npm) Name() string { return pm.name }

func (pm Npm) GetOutdated() ([]dependency.Dependency, error) {
	output, err := exec.Command("npm", "outdated", "--json", "--long").Output()

	// npm outdated returns exit status 1 if there are outdated packages
	if err != nil && err.Error() != "exit status 1" {
		return nil, err
	}

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	for _, d := range outdated {
		err = validate.Struct(d)
		if err != nil {
			return nil, err
		}
	}

	outdatedDeps := make([]dependency.Dependency, len(outdated))

	var index int
	for name, value := range outdated {
		d, err := dependency.New(
			name,
			value.Wanted,
			value.Latest,
			value.Current,
			value.Homepage,
		)
		if err != nil {
			return nil, fmt.Errorf("invalid package versions: %v %w", value, err)
		}

		outdatedDeps[index] = *d
		index++
	}

	return outdatedDeps, nil
}

func (pm Npm) Update(deps []dependency.Dependency) ([]byte, error) {
	args := make([]string, len(deps)+1)
	args[0] = "install"
	for i, d := range deps {
		args[i+1] = d.Name
		if d.Target == dependency.Latest {
			args[i+1] += "@latest"
		}
	}
	slog.Info("install args:", "args", args)

	output, err := exec.Command("npm", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (npm Npm) String() string {
	return npm.Name()
}
