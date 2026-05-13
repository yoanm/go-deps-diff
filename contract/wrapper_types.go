package contract

// PackageMap holds package information for efficient lookup.
//
// Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.
// PackageMap is used as input to depsdiff.Diff to represent the state of dependencies at a point in time
// (typically parsed from a package lock file).
type PackageMap map[string]PkgWrapper

// PkgWrapper is an interface that provides access to package metadata.
//
// Implementations should provide information about a package from a lock file,
// including version, dependency type (regular/dev), and status information.
type PkgWrapper interface {
	// GetName returns the package name (e.g., "vendor/package", "symfony/console").
	GetName() string
	// GetVersion returns the package version information.
	// For removed packages in a comparison, this returns the last known version.
	GetVersion() PkgVersion
	// IsAbandoned returns true if the package is marked as abandoned (no longer maintained).
	IsAbandoned() bool
	// IsDevOnly returns true if the package is only for development environment (dev-only dependency).
	// A package may exist only as a root dev requirement from user perspective, but actually used
	// by a non-dev requirement.
	IsDevOnly() bool
	// IsRootRequirement returns true if the package is explicitly required (typically from a requirement/manifest file).
	IsRootRequirement() bool
	// IsRootDevRequirement returns true if the package is explicitly required for development only
	// (e.g., in "require-dev" section for composer or "devDependencies" for npm).
	IsRootDevRequirement() bool
	// GetLink returns the best available link for the package (wiki, docs, source, homepage, etc.)
	// or an empty string if no link is available.
	GetLink() string
}

// PkgVersion contains version information for a package.
//
// Raw contains the exact version string as defined in the lock file.
// Label is the human-readable version representation.
// Semver is populated only if Raw is a valid semantic version, otherwise nil.
type PkgVersion struct {
	// Raw is the value as defined in the lock file (e.g., "1.2.3", "dev-master#abcd123", "1.0.0-beta").
	Raw string
	// Label is the human-readable title for the version (e.g., "1.2.3", "dev-master#abcd123").
	Label string
	// Semver will be defined only if Raw version is semver compliant, otherwise it will be nil.
	// Check if this is non-nil before using semantic version components.
	Semver *Semver
}
