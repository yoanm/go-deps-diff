# contract

Package contract defines the core data types and interfaces for dependency comparison.

This package provides:

```diff
- PackageMap: A map for efficient O(1) lookup of package information
- PackageChange: Detailed information about how a package changed
- PkgWrapper and PkgVersion: Abstractions for package data from lock files
- Operation types: Semantic representation of what changed (added, removed, upgraded, etc.)
- Semver: Semantic version components (major, minor, patch, extra)
```

Types in this package are used as inputs to and outputs from the depsdiff.Diff function.
They define the contract between package managers and the diff engine.

## Types

### type [DiffMap](./diff_types.go#L6)

`type DiffMap map[string]*PackageChange`

DiffMap maps package names to information about what changed.
Key is the package name (e.g., "vendor/package"), value is a pointer to PackageChange
containing the operation type and version information.

### type [Operation](./operation_types.go#L49)

`type Operation struct { ... }`

Operation describes what changed about a package: the operation name and semantic version type.

Name indicates the type of change (addition, removal, upgrade, downgrade, etc.).
SemverType indicates which semantic version component changed (for updated packages).

### type [OperationName](./operation_types.go#L4)

`type OperationName string`

OperationName describes what changed about a package (addition, removal, update direction).

#### Constants

```golang
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
```

### type [OperationSemverType](./operation_types.go#L27)

`type OperationSemverType string`

OperationSemverType describes the semantic version component type for updated packages.

This indicates which semantic version component changed (e.g., MAJOR, MINOR, PATCH).
It is only relevant for updated packages; for added and removed packages,
the SemverType is SemverNoUpdate since only one version is available.

#### Constants

```golang
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
```

### type [PackageChange](./diff_types.go#L18)

`type PackageChange struct { ... }`

PackageChange contains detailed information about a package difference.

The Package field holds a reference to the package wrapper (agnostic of the package manager).
See PkgWrapper for more information.

The Operation field indicates what changed (ADDITION, REMOVAL, UPGRADE, DOWNGRADE, etc.)
and the semantic version type of the change (MAJOR, MINOR, PATCH, EXTRA, UNKNOWN, NONE).

PreviousVersion is only populated for updated packages (when Operation is UPGRADE or DOWNGRADE).
For added and removed packages, PreviousVersion is empty (zero value).

### type [PackageMap](./wrapper_types.go#L8)

`type PackageMap map[string]PkgWrapper`

PackageMap holds package information for efficient lookup.

Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.
PackageMap is used as input to depsdiff.Diff to represent the state of dependencies at a point in time
(typically parsed from a package lock file).

### type [PkgVersion](./wrapper_types.go#L41)

`type PkgVersion struct { ... }`

PkgVersion contains version information for a package.

Raw contains the exact version string as defined in the lock file.
Label is the human-readable version representation.
Semver is populated only if Raw is a valid semantic version, otherwise nil.

### type [PkgWrapper](./wrapper_types.go#L14)

`type PkgWrapper interface { ... }`

PkgWrapper is an interface that provides access to package metadata.

Implementations should provide information about a package from a lock file,
including version, dependency type (regular/dev), and status information.

### type [Semver](./semver.go#L10)

`type Semver struct { ... }`

Semver represents semantic version components.

A semantic version follows the format MAJOR.MINOR.PATCH[+extra].
For example, version "1.2.3-beta" has Major=1, Minor=2, Patch=3, Extra="-beta".

This struct is only populated when a version string is determined to be
valid semantic version format. See contract.PkgVersion.Semver.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
