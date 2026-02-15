package dependency

import (
	"fmt"

	"updep/pkg/version"
)

type Target int

const (
	None Target = iota
	Wanted
	Latest
)

type Dependency struct {
	Name     string
	Wanted   version.Version
	Latest   version.Version
	Current  version.Version
	Target   Target
	Homepage string
}

func New(
	name string,
	wantedVersion string,
	latestVersion string,
	currentVersion string,
	homepage string,
) (*Dependency, error) {
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

	return &Dependency{
		Name:     name,
		Wanted:   *wanted,
		Latest:   *latest,
		Current:  *current,
		Target:   None,
		Homepage: homepage,
	}, nil
}

func (d *Dependency) ToggleTarget() {
	if d.Target != None {
		d.Target = None
		return
	}

	d.SelectTarget()
}

func (d *Dependency) SelectTarget() {
	if d.Target != None {
		return
	}

	if d.Current.Compare(d.Wanted) >= 0 {
		d.Target = Latest
	} else {
		d.Target = Wanted
	}
}

func (d *Dependency) DiffLevel() version.DiffLevel {
	switch d.Target {
	case Wanted:
		return version.VersionDiffLevel(d.Current, d.Wanted)
	case Latest:
		return version.VersionDiffLevel(d.Current, d.Latest)
	default:
		return version.None
	}
}
