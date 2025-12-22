package packages

import (
	"errors"
)

type Package struct {
	Name    string
	Wanted  Version
	Latest  Version
	Current Version
}

func NewPackage(
	name string,
	wantedVersion string,
	latestVersion string,
	currentVersion string,
) (*Package, error) {
	wanted, err := NewVersion(wantedVersion)
	if err != nil {
		return nil, errors.New("invalid packagemanager")
	}
	latest, err := NewVersion(latestVersion)
	if err != nil {
		return nil, errors.New("invalid packagemanager")
	}
	current, err := NewVersion(currentVersion)
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
