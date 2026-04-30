package depsdiff

// Operation describes what changed (operation, semver type) about a package

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
	// SemverMajorUpdate is for updated packages where the major component differs.
	SemverMajorUpdate OperationSemverType = "MAJOR"
	// SemverMinorUpdate is for updated packages where the minor component differs.
	SemverMinorUpdate OperationSemverType = "MINOR"
	// SemverPatchUpdate is for updated packages where the patch component differs.
	SemverPatchUpdate OperationSemverType = "PATCH"
	// SemverExtraUpdate is for updated packages where the extra component differs.
	SemverExtraUpdate OperationSemverType = "EXTRA"
	// SemverUnknownUpdate is for updated packages where we can't determine the difference (e.g., non-semver versions).
	SemverUnknownUpdate OperationSemverType = "UNKNOWN"
	// SemverNoUpdate is for added and removed packages (=no difference as only one version available).
	SemverNoUpdate OperationSemverType = "NONE"
)

type Operation struct {
	Name       OperationName
	SemverType OperationSemverType // Semver[...]Update
}
