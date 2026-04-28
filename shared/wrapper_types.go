package shared

// PackageMap holds package information for efficient lookup
// Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.
type PackageMap map[string]PkgWrapper

type PkgWrapper interface {
	GetName() string // Returns the package name (e.g., "vendor/package")
	// GetVersion returns the package Version (e.g., "1.2.3" or "dev-master").
	// For removed packages, this returns the last known Version.
	GetVersion() PkgVersion
	IsAbandoned() bool // true if the package is marked as abandoned (no longer maintained)
	// IsDevOnly is true if package is only for dev environment (dev-only dependency).
	// A package may exist only as root dev requirement (from user point of view), but actually used
	// by a non-dev requirement !
	IsDevOnly() bool
	IsRootRequirement() bool // true if package is explicitly required (requirement file usually)
	// IsRootDevRequirement is true if package is explicitly required (requirement file usually),
	// but only for dev environment
	IsRootDevRequirement() bool
	// GetLink Returns the best available link for the package (wiki, docs, source, homepage, etc.)
	// or empty string if none available
	GetLink() string
}

type PkgVersion struct {
	Raw   string // Raw value as defined in the lock file
	Label string // Human-readable label (e.g., "1.2.3", "dev-master#abcd123")
}
