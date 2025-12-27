package packages

import (
	"fmt"

	"updep/pkg/version"
)

type Package struct {
	Name    string
	Wanted  version.Version
	Latest  version.Version
	Current version.Version
	// Target is the version selected for upgrade/downgrade, or nil if not yet chosen
	Target   *version.Version
	Homepage string
}

func New(
	name string,
	wantedVersion string,
	latestVersion string,
	currentVersion string,
	homepage string,
) (*Package, error) {
	wanted, err := version.New(wantedVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid wanted version: %w", err)
	}
	latest, err := version.New(latestVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid latest version: %w", err)
	}
	current, err := version.New(currentVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid current version: %w", err)
	}

	return &Package{
		Name:     name,
		Wanted:   *wanted,
		Latest:   *latest,
		Current:  *current,
		Target:   nil,
		Homepage: homepage,
	}, nil
}
