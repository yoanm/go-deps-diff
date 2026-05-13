package contract

// OperationName describes what changed about a package (addition, removal, update direction).
type OperationName string

const (
	// AdditionOperation indicates a package was added (exists in current but not in previous).
	AdditionOperation OperationName = "ADDITION"
	// RemovalOperation indicates a package was removed (exists in previous but not in current).
	RemovalOperation OperationName = "REMOVAL"
	// UpgradeOperation indicates a package version was increased.
	UpgradeOperation OperationName = "UPGRADE"
	// DowngradeOperation indicates a package version was decreased.
	DowngradeOperation OperationName = "DOWNGRADE"
	// UnknownUpdateOperation indicates a package was updated but direction is unknown
	// (typically for non-semver versions where comparison is not possible).
	UnknownUpdateOperation OperationName = "UNKNOWN_UPDATE"
	// NoChangeOperation indicates a package version has not changed.
	NoChangeOperation OperationName = "NONE"
)

// OperationSemverType describes the semantic version component type for updated packages.
//
// This indicates which semantic version component changed (e.g., MAJOR, MINOR, PATCH).
// It is only relevant for updated packages; for added and removed packages,
// the SemverType is SemverNoUpdate since only one version is available.
type OperationSemverType string

const (
	// SemverMajorUpdate indicates the major version component differs between versions.
	SemverMajorUpdate OperationSemverType = "MAJOR"
	// SemverMinorUpdate indicates the minor version component differs between versions.
	SemverMinorUpdate OperationSemverType = "MINOR"
	// SemverPatchUpdate indicates the patch version component differs between versions.
	SemverPatchUpdate OperationSemverType = "PATCH"
	// SemverExtraUpdate indicates the extra/pre-release component differs between versions
	// (major, minor, patch are equal but extra differs).
	SemverExtraUpdate OperationSemverType = "EXTRA"
	// SemverUnknownUpdate indicates the difference cannot be determined (e.g., one or both versions are non-semver).
	SemverUnknownUpdate OperationSemverType = "UNKNOWN"
	// SemverNoUpdate indicates no semantic version difference (used for added, removed, or unchanged packages).
	SemverNoUpdate OperationSemverType = "NONE"
)

// Operation describes what changed about a package: the operation name and semantic version type.
//
// Name indicates the type of change (addition, removal, upgrade, downgrade, etc.).
// SemverType indicates which semantic version component changed (for updated packages).
type Operation struct {
	// Name describes the operation type (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, UNKNOWN_UPDATE, NONE).
	Name OperationName
	// SemverType describes the semantic version component type for updated packages
	// (MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, NONE). Only relevant for updated packages.
	SemverType OperationSemverType
}
