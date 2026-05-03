package depsdiff

import (
	"github.com/yoanm/go-deps-diff/shared"
)

// Diff compares two packages maps and returns the differences.
func Diff(previous, current shared.PackageMap) (shared.DiffMap, error) {
	// Find differences
	output := shared.DiffMap{}

	// Find added and updated packages
	for name, currentPkg := range current {
		pkgChange := shared.PackageChange{ //nolint:exhaustruct // Other properties will be filled based on the operation
			Package: currentPkg,
		}

		if previousPkg, previousExists := previous[name]; previousExists {
			previousVersion := previousPkg.GetVersion().Raw
			currentVersion := currentPkg.GetVersion().Raw

			if previousVersion != currentVersion {
				pkgChange.Operation = guessUpdateOperation(previousVersion, currentVersion)
				pkgChange.PreviousVersion = previousPkg.GetVersion()
			} else {
				pkgChange.Operation = shared.Operation{Name: shared.NoChangeOperation, SemverType: shared.SemverNoUpdate}
			}
		} else {
			pkgChange.Operation = shared.Operation{Name: shared.AdditionOperation, SemverType: shared.SemverNoUpdate}
		}

		output[name] = &pkgChange
	}

	// Find removed packages
	for name, previousPkg := range previous {
		if _, exists := current[name]; !exists {
			info := shared.PackageChange{ //nolint:exhaustruct // PreviousVersion is unused for removed packages !
				Package:   previousPkg,
				Operation: shared.Operation{Name: shared.RemovalOperation, SemverType: shared.SemverNoUpdate},
			}

			output[name] = &info
		}
	}

	return output, nil
}

// guessUpdateOperation detects the type and direction of a version update.
func guessUpdateOperation(previousVersion, currentVersion string) shared.Operation {
	result := shared.Operation{
		Name:       shared.UnknownUpdateOperation,
		SemverType: shared.SemverUnknownUpdate,
	}

	prevTag, invalidPrevErr := shared.ParseSemverVersion(previousVersion)
	currTag, invalidCurrentErr := shared.ParseSemverVersion(currentVersion)

	// If either version is not semver, direction is UNKNOWN
	if invalidPrevErr != nil || invalidCurrentErr != nil {
		return result
	}

	switch {
	case prevTag.Major != currTag.Major:
		result.SemverType = shared.SemverMajorUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Major, currTag.Major)
	case prevTag.Minor != currTag.Minor:
		result.SemverType = shared.SemverMinorUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Minor, currTag.Minor)
	case prevTag.Patch != currTag.Patch:
		result.SemverType = shared.SemverPatchUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Patch, currTag.Patch)
	// All numeric components equal -> Compare extra components (pre-release or build metadata)
	case prevTag.Extra != currTag.Extra:
		result.SemverType = shared.SemverExtraUpdate
		result.Name = shared.UnknownUpdateOperation
	}

	return result
}

func guessDirectionFromSemverComponent(prev, curr int) shared.OperationName {
	if curr > prev {
		return shared.UpgradeOperation
	} else if curr < prev {
		return shared.DowngradeOperation
	}

	return shared.UnknownUpdateOperation
}
