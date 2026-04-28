package depsdiff

// Operation describes what changed (operation, direction, semver type) about a package

type OperationName string

const (
	AdditionOperation      OperationName = "ADDITION"
	RemovalOperation       OperationName = "REMOVAL"
	UpgradeOperation       OperationName = "UPGRADE"
	DowngradeOperation     OperationName = "DOWNGRADE"
	UnknownUpdateOperation OperationName = "UNKNOWN_UPDATE"
)

// OperationSemverType describes the type of the change for updated packages (e.g., whether it's a major, minor, patch,
// extra, unknown or none change).
// It's not relevant for added and removed packages, which are considered as "NONE" (no semver difference).
type OperationSemverType string

const (
	// DiffSemverMajor is for updated packages where the major component differs.
	DiffSemverMajor OperationSemverType = "MAJOR"
	// DiffSemverMinor is for updated packages where the minor component differs.
	DiffSemverMinor OperationSemverType = "MINOR"
	// DiffSemverPatch is for updated packages where the patch component differs.
	DiffSemverPatch OperationSemverType = "PATCH"
	// DiffSemverExtra is for updated packages where the extra component differs.
	DiffSemverExtra OperationSemverType = "EXTRA"
	// DiffSemverUnknown is for updated packages where we can't determine the difference (e.g., non-semver versions).
	DiffSemverUnknown OperationSemverType = "UNKNOWN"
	// DiffSemverNone is for added and removed packages (=no difference as only one version available).
	DiffSemverNone OperationSemverType = "NONE"
)

type Operation struct {
	Name       OperationName
	SemverType OperationSemverType // DIFF_SEMVER_*
}
