package depsdiff

import "github.com/yoanm/go-deps-diff/shared"

type Input struct {
	Current  PkgManagerInput
	Previous PkgManagerInput
}
type PkgManagerInput struct {
	// Lock represents the content of the lock file (e.g., composer.lock for composer, package-lock.json for npm,
	// yarn.lock for yarn, etc...)
	Lock []byte
	// Requirement (Optional) represents the content of the requirement file  (e.g. composer.json for composer,
	// package.json for npm/yarn, etc...). This is optional and may be empty if not available.
	// It can be used to provide additional context about the packages (e.g., whether they are dev requirement or not).
	Requirement []byte
}

// Output is the result of comparing two composer.lock files.
type Output struct {
	Changes map[string]PackageChange
}

// PackageChange contains detailed information about a package difference.
type PackageChange struct {
	Package   shared.PkgWrapper
	Operation Operation

	PreviousVersion shared.PkgVersion // Only available for updated packages ! Empty (zero value) otherwise.
}
