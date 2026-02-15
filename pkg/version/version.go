package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Suffix string
}

func New(str string) (*Version, error) {
	versions := strings.Split(str, ".")

	if len(versions) < 3 {
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

	patchStr, suffix, _ := strings.Cut(versions[2], "-")
	patch, err := strconv.Atoi(patchStr)
	if err != nil {
		return nil, errors.New("invalid version")
	}

	return &Version{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Suffix: suffix,
	}, nil
}

func (v Version) String() string {
	var suffix string
	if v.Suffix != "" {
		suffix = "-" + v.Suffix
	}

	return fmt.Sprintf(
		"%d.%d.%d%s",
		v.Major,
		v.Minor,
		v.Patch,
		suffix,
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
