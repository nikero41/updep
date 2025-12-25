package packages

import (
	"errors"

	"updep/pkg/version"
)

type Package struct {
	Name    string
	Wanted  version.Version
	Latest  version.Version
	Current version.Version
}

func NewPackage(
	name string,
	wantedVersion string,
	latestVersion string,
	currentVersion string,
) (*Package, error) {
	wanted, err := version.New(wantedVersion)
	if err != nil {
		return nil, errors.New("invalid packagemanager")
	}
	latest, err := version.New(latestVersion)
	if err != nil {
		return nil, errors.New("invalid packagemanager")
	}
	current, err := version.New(currentVersion)
	if err != nil {
		return nil, errors.New("invalid packagemanager")
	}

	return &Package{
		Name:    name,
		Wanted:  *wanted,
		Latest:  *latest,
		Current: *current,
	}, nil
}
