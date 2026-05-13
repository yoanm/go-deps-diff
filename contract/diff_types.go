package contract

// DiffMap maps package names to information about what changed.
// Key is the package name (e.g., "vendor/package"), value is a pointer to PackageChange
// containing the operation type and version information.
type DiffMap map[string]*PackageChange

// PackageChange contains detailed information about a package difference.
//
// The Operation field indicates what changed (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, etc.)
// and the semantic version type of the change (MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, NONE).
//
// PreviousVersion is only populated for updated packages (when Operation is UPGRADE or DOWNGRADE).
// For added and removed packages, PreviousVersion is empty (zero value).
type PackageChange struct {
	Package   PkgWrapper
	Operation Operation

	PreviousVersion PkgVersion // Only available for updated packages ! Empty (zero value) otherwise.
}
