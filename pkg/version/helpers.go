package version

type DiffLevel int

const (
	Major DiffLevel = iota
	Minor
	Patch
	None
)

func VersionDiffLevel(a Version, b Version) DiffLevel {
	if b.Major != a.Major {
		return Major
	} else if b.Minor != a.Minor {
		return Minor
	} else if b.Patch != a.Patch {
		return Patch
	}

	return None
}
