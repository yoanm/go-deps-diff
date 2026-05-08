package contract

// PackageMap holds package information for efficient lookup
// Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.
type PackageMap map[string]PkgWrapper

type PkgWrapper interface {
	GetName() string // Returns the package name (e.g., "vendor/package")
	// GetVersion returns the package Version (e.g., "1.2.3" or "dev-master").
	// For removed packages, this returns the last known Version.
	GetVersion() PkgVersion
	// IsAbandoned is true if the package is marked as abandoned (no longer maintained)
	IsAbandoned() bool
	// IsDevOnly is true if package is only for dev environment (dev-only dependency).
	// A package may exist only as root dev requirement (from user point of view), but actually used
	// by a non-dev requirement !
	IsDevOnly() bool
	// IsRootRequirement is true if package is explicitly required (requirement file usually)
	IsRootRequirement() bool
	// IsRootDevRequirement is true if package is explicitly required (requirement file usually),
	// but only for dev environment (e.g. "require-dev" section for composer)
	IsRootDevRequirement() bool
	// GetLink Returns the best available link for the package (wiki, docs, source, homepage, etc.)
	// or empty string if none available
	GetLink() string
}

type PkgVersion struct {
	// Raw is the value as defined in the lock file
	Raw string
	// Label is the Human-readable title for the version (e.g., "1.2.3", "dev-master#abcd123")
	Label string
	// Semver will be defined only if Raw version is semver compliant, otherwise it will be nil
	Semver *Semver
}
