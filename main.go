// Package depsdiff provides functionality to compare two dependency package maps and identify differences.
//
// The main API is the Diff function, which compares two PackageMaps (representing lock file states
// at different points in time) and returns a DiffMap containing detailed information about what changed:
// added, removed, and updated packages with semantic version analysis.
//
// Example usage:
//
//	previousMap := contract.PackageMap{/* packages from old lock file */}
//	currentMap := contract.PackageMap{/* packages from new lock file */}
//
//	for pkgName, change := range depsdiff.Diff(previousMap, currentMap) {
//		fmt.Printf("%s: %s\n", pkgName, change.Operation.Name)
//	}
package depsdiff

import (
	"github.com/yoanm/go-deps-diff/contract"
)

// Diff compares two package maps and returns detailed information about differences.
//
// Parameters:
//   - previous: PackageMap representing the previous state (e.g., packages from old lock file)
//   - current: PackageMap representing the current state (e.g., packages from new lock file)
//
// Returns:
//   - DiffMap: A map where keys are package names and values contain PackageChange information
//     including the operation type (added, removed, upgraded, downgraded, etc.) and semantic
//     version analysis for updated packages.
//
// For each package in the diff result:
//   - PackageChange.Package field holds a reference to the package wrapper (agnostic of the package manager).
//     See contract.PkgWrapper for more information.
//   - PackageChange.Operation indicates what changed (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, etc.)
//   - PackageChange.Operation.SemverType indicates the type of change (MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, NONE)
//   - PackageChange.PreviousVersion is only populated for updated packages
func Diff(previous, current contract.PackageMap) contract.DiffMap {
	output := contract.DiffMap{}

	// Find added and updated packages
	for name, currentPkg := range current {
		pkgChange := contract.PackageChange{ //nolint:exhaustruct // Other properties will be filled based on the operation
			Package: currentPkg,
		}

		if previousPkg, previousExists := previous[name]; previousExists {
			if previousPkg.GetVersion().Raw != currentPkg.GetVersion().Raw {
				pkgChange.Operation = guessUpdateOperation(previousPkg.GetVersion(), currentPkg.GetVersion())
				pkgChange.PreviousVersion = previousPkg.GetVersion()
			} else {
				pkgChange.Operation = contract.Operation{Name: contract.NoChangeOperation, SemverType: contract.SemverNoUpdate}
			}
		} else {
			pkgChange.Operation = contract.Operation{Name: contract.AdditionOperation, SemverType: contract.SemverNoUpdate}
		}

		output[name] = &pkgChange
	}

	// Find removed packages
	for name, previousPkg := range previous {
		if _, exists := current[name]; !exists {
			output[name] = &contract.PackageChange{ //nolint:exhaustruct // PreviousVersion is unused for removed packages !
				Package:   previousPkg,
				Operation: contract.Operation{Name: contract.RemovalOperation, SemverType: contract.SemverNoUpdate},
			}
		}
	}

	return output
}

// guessUpdateOperation detects the type and direction of a version update.
func guessUpdateOperation(previous, current contract.PkgVersion) contract.Operation {
	result := contract.Operation{
		Name:       contract.UnknownUpdateOperation,
		SemverType: contract.SemverUnknownUpdate,
	}

	// If either version is not semver, direction is UNKNOWN
	if previous.Semver == nil || current.Semver == nil {
		return result
	}

	switch {
	case previous.Semver.Major != current.Semver.Major:
		result.Name = guessDirectionFromSemverComponent(previous.Semver.Major, current.Semver.Major)
		result.SemverType = contract.SemverMajorUpdate
	case previous.Semver.Minor != current.Semver.Minor:
		result.Name = guessDirectionFromSemverComponent(previous.Semver.Minor, current.Semver.Minor)
		result.SemverType = contract.SemverMinorUpdate
	case previous.Semver.Patch != current.Semver.Patch:
		result.Name = guessDirectionFromSemverComponent(previous.Semver.Patch, current.Semver.Patch)
		result.SemverType = contract.SemverPatchUpdate
	// All numeric components equal -> Compare extra components (pre-release or build metadata)
	case previous.Semver.Extra != current.Semver.Extra:
		result.SemverType = contract.SemverExtraUpdate
	}

	return result
}

func guessDirectionFromSemverComponent(prev, curr int) contract.OperationName {
	if curr > prev {
		return contract.UpgradeOperation
	} else if curr < prev {
		return contract.DowngradeOperation
	}

	return contract.UnknownUpdateOperation
}
