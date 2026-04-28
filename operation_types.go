package depsdiff

// Operation describes what changed (operation, direction, semver type) about a package

type OperationName string

const (
	AddedPackage    OperationName = "ADDITION"
	RemovedPackaged OperationName = "REMOVAL"
	UpdatedPackage  OperationName = "UPDATE"
)

type OperationSemverType string

const (
	DiffSemverMajor   OperationSemverType = "MAJOR"
	DiffSemverMinor   OperationSemverType = "MINOR"
	DiffSemverPatch   OperationSemverType = "PATCH"
	DiffSemverExtra   OperationSemverType = "EXTRA"
	DiffSemverUnknown OperationSemverType = "UNKNOWN"
	DiffSemverNone    OperationSemverType = "NONE"
)

type OperationDirection string

const (
	DiffDirectionUp      OperationDirection = "UP"
	DiffDirectionDown    OperationDirection = "DOWN"
	DiffDirectionUnknown OperationDirection = "UNKNOWN"
	DiffDirectionNone    OperationDirection = "NONE"
)

type Operation struct {
	Name       OperationName
	SemverType OperationSemverType // Only for Name=UpdatedPackage: DIFF_SEMVER_*
	Direction  OperationDirection  // Only for Name=UpdatedPackage: DIFF_DIRECTION_*
}
