package depsdiff

import (
	"github.com/yoanm/go-deps-diff/contract"
	"time"
)

// Diff compares two packages maps and returns the differences.
func Diff(previous, current contract.PackageMap) (contract.DiffMap, error) {
	time.Sleep(1 * time.Millisecond)
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

	return output, nil
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
