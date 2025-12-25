package version

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

func New(str string) (*Version, error) {
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

func (v Version) RenderWithDiff(b Version) string {
	switch VersionDiffLevel(v, b) {
	case Major:
		return majorDiffStyle.Render(v.String())

	case Minor:
		return fmt.Sprintf(
			"%d.%v",
			v.Major,
			minorDiffStyle.Render(fmt.Sprintf("%d.%d", v.Minor, v.Patch)),
		)

	case Patch:
		return fmt.Sprintf(
			"%d.%d.%v",
			v.Major,
			v.Minor,
			patchDiffStyle.Render(fmt.Sprintf("%d", v.Patch)),
		)

	default:
		return v.String()
	}
}
