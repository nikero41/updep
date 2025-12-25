package version

type DiffLevel int

const (
	Major DiffLevel = iota
	Minor
	Patch
	None
)

func VersionDiffLevel(v Version, b Version) DiffLevel {
	diff := Version{
		Major: b.Major - v.Major,
		Minor: b.Minor - v.Minor,
		Patch: b.Patch - v.Patch,
	}

	if diff.Major != 0 {
		return Major
	} else if diff.Minor != 0 {
		return Minor
	} else if diff.Patch != 0 {
		return Patch
	}

	return None
}
