package diff

// Output is the result of comparing two composer.lock files
type Output struct {
	Packages []PackageInfo // @TODO: MAJOR REFACTOR - Change to map[string]PackageInfo (key=package name). Requires updating: analyzer.go (3 append ops), all tests (20+ assertions), README examples
}

// PackageInfo contains detailed information about a package difference
type PackageInfo struct {
	Name                 string
	IsRootRequirement    bool
	IsRootDevRequirement bool
	IsAbandoned          bool
	Link                 string
	Update               UpdateType
	Previous             PkgVersion
	Current              PkgVersion
}

// UpdateType describes what changed about a package
type UpdateType struct {
	Type      string // 'ADDED', 'REMOVED', 'UPDATED'
	SubType   string // Only for Type='UPDATED': 'MAJOR', 'MINOR', 'PATCH'. For other types: 'NONE'
	Direction string // Only for Type='UPDATED': 'UP', 'DOWN', 'NONE', 'UNKNOWN'. For other types: 'NONE'
}

// PkgVersion interface allows for version discrimination
type PkgVersion interface {
	GetFull() string
	GetType() string
}

// PkgVersionTag represents a semver-style version
type PkgVersionTag struct {
	Full  string // Full version string from composer.lock (e.g., "1.2.3", "v2.0.0-beta")
	Type  string // 'TAG' (always for this type)
	Major string // Semver major version
	Minor string // Semver minor version
	Patch string // Semver patch version
	Extra string // Optional pre-release/build metadata (e.g., "-beta.1", "+build123")
}

func (v *PkgVersionTag) GetFull() string {
	return v.Full
}

func (v *PkgVersionTag) GetType() string {
	return v.Type
}

// PkgVersionCommit represents a commit-based version
type PkgVersionCommit struct {
	Full   string // Full version string (commit hash)
	Type   string // 'COMMIT' (always for this type)
	Commit string // Commit hash
}

func (v *PkgVersionCommit) GetFull() string {
	return v.Full
}

func (v *PkgVersionCommit) GetType() string {
	return v.Type
}

// NewPkgVersionTag creates a new PkgVersionTag
func NewPkgVersionTag(version, major, minor, patch, extra string) *PkgVersionTag {
	return &PkgVersionTag{
		Full:  version,
		Type:  "TAG",
		Major: major,
		Minor: minor,
		Patch: patch,
		Extra: extra,
	}
}

// NewPkgVersionCommit creates a new PkgVersionCommit
func NewPkgVersionCommit(version, commit string) *PkgVersionCommit {
	full := version
	// If version differs from commit, append short commit hash (first 7 chars)
	if version != commit && len(commit) >= 7 {
		full = version + "#" + commit[:7]
	}
	return &PkgVersionCommit{
		Full:   full,
		Type:   "COMMIT",
		Commit: commit,
	}
}
