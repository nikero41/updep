package packagemanager

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// NOTE: Version will break currently if version has prefix or suffix

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(str string) (*Version, error) {
	versions := strings.Split(str, ".")

	if len(versions) != 3 {
		return nil, errors.New("invalid version")
	}

	major, err := strconv.Atoi(versions[0])
	if err != nil {
		return nil, errors.New("invalid version")
	}
	minor, err := strconv.Atoi(versions[1])
	if err != nil {
		return nil, errors.New("invalid version")
	}
	patch, err := strconv.Atoi(versions[2])
	if err != nil {
		return nil, errors.New("invalid version")
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func (v Version) String() string {
	return fmt.Sprintf(
		"%d.%d.%d",
		v.Major,
		v.Minor,
		v.Patch,
	)
}

func (v Version) Compare(b Version) int {
	if v.Major > b.Major {
		return 1
	} else if v.Major < b.Major {
		return -1
	}

	if v.Minor > b.Minor {
		return 1
	} else if v.Minor < b.Minor {
		return -1
	}

	if v.Patch > b.Patch {
		return 1
	} else if v.Patch < b.Patch {
		return -1
	}

	return 0
}

func (v Version) Diff(b Version) Version {
	diff := Version{
		Major: b.Major - v.Major,
		Minor: b.Minor - v.Minor,
		Patch: b.Patch - v.Patch,
	}

	return diff
}
