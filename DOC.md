# depsdiff

## Types

### type [Input](/types.go#L5)
<sdf
`type Input struct { ... }`

### type [Operation](/operation_types.go#L35)

`type Operation struct { ... }`

### type [OperationName](/operation_types.go#L5)

`type OperationName string`

#### Constants

```golang
const (
    AdditionOperation      OperationName = "ADDITION"
    RemovalOperation       OperationName = "REMOVAL"
    UpgradeOperation       OperationName = "UPGRADE"
    DowngradeOperation     OperationName = "DOWNGRADE"
    UnknownUpdateOperation OperationName = "UNKNOWN_UPDATE"
)
```

### type [OperationSemverType](/operation_types.go#L18)

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
    // SemverNoUpdate is for added and removed packages (=no difference as only one version available).
    SemverNoUpdate OperationSemverType = "NONE"
)
```

### type [Output](/types.go#L20)

`type Output struct { ... }`

Output is the result of comparing two packages maps.

#### func [ComposerDiff](/manager.go#L9)

`func ComposerDiff(input *Input) (*Output, error)`

#### func [Diff](/analyzer.go#L8)

`func Diff(previous, current shared.PackageMap) (*Output, error)`

Diff compares two packages maps and returns the differences.

### type [PackageChange](/types.go#L25)

`type PackageChange struct { ... }`

PackageChange contains detailed information about a package difference.

### type [PkgManagerInput](/types.go#L9)

`type PkgManagerInput struct { ... }`

## Sub Packages

* [.tools](./.tools)

* [composer](./composer)

* [shared](./shared)

* [shared_test](./shared_test)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
