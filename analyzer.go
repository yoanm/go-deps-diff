package depsdiff

import (
	"github.com/yoanm/go-deps-diff/shared"
)

// Diff compares two composer.lock files and returns the differences
// Note: Currently handles both lock file comparison AND requirement file integration.
func Diff(previous, current shared.PackageMap) (*Output, error) {
	// Find differences
	output := &Output{
		Changes: make(map[string]PackageChange),
	}

	// Find added and updated packages
	for name, currentPkg := range current {
		info := PackageChange{Package: currentPkg} //nolint:exhaustruct // Additional properties will be added below

		if previousPkg, previousExists := previous[name]; previousExists {
			previousVersion := previousPkg.GetVersion().Raw
			currentVersion := currentPkg.GetVersion().Raw

			if previousVersion != currentVersion {
				info.PreviousVersion = previousPkg.GetVersion()

				info.Operation = guessUpdateOperation(previousVersion, currentVersion)

				output.Changes[name] = info
			}
		} else {
			info.Operation = Operation{Name: AddedPackage, SemverType: DiffSemverNone, Direction: DiffDirectionNone}

			output.Changes[name] = info
		}
	}

	// Find removed packages
	for name, previousPkg := range previous {
		if _, exists := current[name]; !exists {
			info := PackageChange{ //nolint:exhaustruct // PreviousVersion is unused for removed packages !
				Package:   previousPkg,
				Operation: Operation{Name: RemovedPackaged, SemverType: DiffSemverNone, Direction: DiffDirectionNone},
			}

			output.Changes[name] = info
		}
	}

	return output, nil
}

// guessUpdateOperation detects the type and direction of a version update.
func guessUpdateOperation(previousVersion, currentVersion string) Operation {
	result := Operation{
		Name:       UpdatedPackage,
		SemverType: DiffSemverUnknown,
		Direction:  DiffDirectionUnknown,
	}

	// If either version is not semver, direction is UNKNOWN
	prevTag, invalidPrevErr := shared.ParseSemverVersion(previousVersion)
	currTag, invalidCurrentErr := shared.ParseSemverVersion(currentVersion)

	if invalidPrevErr != nil || invalidCurrentErr != nil {
		return result
	}

	// Compare MAJOR
	if prevTag.Major != currTag.Major {
		result.SemverType = DiffSemverMajor
		if currTag.Major > prevTag.Major {
			result.Direction = DiffDirectionUp
		} else {
			result.Direction = DiffDirectionDown
		}

		return result
	}

	// Compare MINOR
	if prevTag.Minor != currTag.Minor {
		result.SemverType = DiffSemverMinor
		if currTag.Minor > prevTag.Minor {
			result.Direction = DiffDirectionUp
		} else {
			result.Direction = DiffDirectionDown
		}

		return result
	}

	// Compare PATCH
	if prevTag.Patch != currTag.Patch {
		result.SemverType = DiffSemverPatch
		if currTag.Patch > prevTag.Patch {
			result.Direction = DiffDirectionUp
		} else {
			result.Direction = DiffDirectionDown
		}

		return result
	}

	// All numeric components equal
	// Compare extra components (pre-release or build metadata)
	if prevTag.Extra != currTag.Extra {
		result.SemverType = DiffSemverExtra
		result.Direction = DiffDirectionUnknown

		return result
	}

	return result
}
