package contract

// Semver represents semantic version components.
//
// A semantic version follows the format MAJOR.MINOR.PATCH[+extra].
// For example, version "1.2.3-beta" has Major=1, Minor=2, Patch=3, Extra="-beta".
//
// This struct is only populated when a version string is determined to be
// valid semantic version format. See contract.PkgVersion.Semver.
type Semver struct {
	Major int    // Major version component (e.g., 1 in "1.2.3")
	Minor int    // Minor version component (e.g., 2 in "1.2.3")
	Patch int    // Patch version component (e.g., 3 in "1.2.3")
	Extra string // Additional version component (e.g., "-beta", "+build" in "1.2.3-beta+build")
}
