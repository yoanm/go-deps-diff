package depsdiff

import (
	"github.com/yoanm/go-deps-diff/shared"
)

// Diff compares two packages maps and returns the differences.
func Diff(previous, current shared.PackageMap) (Output, error) {
	// Find differences
	output := Output{}

	// Find added and updated packages
	for name, currentPkg := range current {
		if previousPkg, previousExists := previous[name]; previousExists {
			previousVersion := previousPkg.GetVersion().Raw
			currentVersion := currentPkg.GetVersion().Raw

			if previousVersion != currentVersion {
				output[name] = PackageChange{
					Package:         currentPkg,
					Operation:       guessUpdateOperation(previousVersion, currentVersion),
					PreviousVersion: previousPkg.GetVersion(),
				}
			}
		} else {
			output[name] = PackageChange{ //nolint:exhaustruct // PreviousVersion is unused for added packages !
				Package:   currentPkg,
				Operation: Operation{Name: AdditionOperation, SemverType: SemverNoUpdate},
			}
		}
	}

	// Find removed packages
	for name, previousPkg := range previous {
		if _, exists := current[name]; !exists {
			info := PackageChange{ //nolint:exhaustruct // PreviousVersion is unused for removed packages !
				Package:   previousPkg,
				Operation: Operation{Name: RemovalOperation, SemverType: SemverNoUpdate},
			}

			output[name] = info
		}
	}

	return output, nil
}

// guessUpdateOperation detects the type and direction of a version update.
func guessUpdateOperation(previousVersion, currentVersion string) Operation {
	result := Operation{
		Name:       UnknownUpdateOperation,
		SemverType: SemverUnknownUpdate,
	}

	prevTag, invalidPrevErr := shared.ParseSemverVersion(previousVersion)
	currTag, invalidCurrentErr := shared.ParseSemverVersion(currentVersion)

	// If either version is not semver, direction is UNKNOWN
	if invalidPrevErr != nil || invalidCurrentErr != nil {
		return result
	}

	switch {
	case prevTag.Major != currTag.Major:
		result.SemverType = SemverMajorUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Major, currTag.Major)
	case prevTag.Minor != currTag.Minor:
		result.SemverType = SemverMinorUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Minor, currTag.Minor)
	case prevTag.Patch != currTag.Patch:
		result.SemverType = SemverPatchUpdate
		result.Name = guessDirectionFromSemverComponent(prevTag.Patch, currTag.Patch)
	// All numeric components equal -> Compare extra components (pre-release or build metadata)
	case prevTag.Extra != currTag.Extra:
		result.SemverType = SemverExtraUpdate
		result.Name = UnknownUpdateOperation
	}

	return result
}

func guessDirectionFromSemverComponent(prev, curr int) OperationName {
	if curr > prev {
		return UpgradeOperation
	} else if curr < prev {
		return DowngradeOperation
	}

	return UnknownUpdateOperation
}
