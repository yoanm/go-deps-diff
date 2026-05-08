# contract

## Types

### type [DiffMap](./diff_types.go#L3)

`type DiffMap map[string]*PackageChange`

### type [Operation](./operation_types.go#L37)

`type Operation struct { ... }`

### type [OperationName](./operation_types.go#L5)

`type OperationName string`

#### Constants

```golang
const (
    AdditionOperation      OperationName = "ADDITION"
    RemovalOperation       OperationName = "REMOVAL"
    UpgradeOperation       OperationName = "UPGRADE"
    DowngradeOperation     OperationName = "DOWNGRADE"
    UnknownUpdateOperation OperationName = "UNKNOWN_UPDATE"
    NoChangeOperation      OperationName = "NONE"
)
```

### type [OperationSemverType](./operation_types.go#L19)

`type OperationSemverType string`

OperationSemverType describes the type of the change for updated packages (e.g., whether it's a major, minor, patch,
extra, unknown or none change).
It's not relevant for added and removed packages, which are considered as "NONE" (no semver difference).

#### Constants

```golang
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
    // SemverNoUpdate is for added and removed packages (=no difference as only one version available)
    // and none operation.
    SemverNoUpdate OperationSemverType = "NONE"
)
```

### type [PackageChange](./diff_types.go#L6)

`type PackageChange struct { ... }`

PackageChange contains detailed information about a package difference.

### type [PackageMap](./wrapper_types.go#L5)

`type PackageMap map[string]PkgWrapper`

PackageMap holds package information for efficient lookup
Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.

### type [PkgVersion](./wrapper_types.go#L28)

`type PkgVersion struct { ... }`

### type [PkgWrapper](./wrapper_types.go#L7)

`type PkgWrapper interface { ... }`

### type [Semver](./semver.go#L3)

`type Semver struct { ... }`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
