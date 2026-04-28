# depsdiff

## Types

### type [Input](/types.go#L5)

`type Input struct { ... }`

### type [Operation](/operation_types.go#L33)

`type Operation struct { ... }`

### type [OperationDirection](/operation_types.go#L24)

`type OperationDirection string`

#### Constants

```golang
const (
    DiffDirectionUp      OperationDirection = "UP"
    DiffDirectionDown    OperationDirection = "DOWN"
    DiffDirectionUnknown OperationDirection = "UNKNOWN"
    DiffDirectionNone    OperationDirection = "NONE"
)
```go

### type [OperationName](/operation_types.go#L5)

`type OperationName string`

#### Constants

```golang
const (
    AddedPackage    OperationName = "ADDITION"
    RemovedPackaged OperationName = "REMOVAL"
    UpdatedPackage  OperationName = "UPDATE"
)
```go

### type [OperationSemverType](/operation_types.go#L13)

`type OperationSemverType string`

#### Constants

```golang
const (
    DiffSemverMajor   OperationSemverType = "MAJOR"
    DiffSemverMinor   OperationSemverType = "MINOR"
    DiffSemverPatch   OperationSemverType = "PATCH"
    DiffSemverExtra   OperationSemverType = "EXTRA"
    DiffSemverUnknown OperationSemverType = "UNKNOWN"
    DiffSemverNone    OperationSemverType = "NONE"
)
```go

### type [Output](/types.go#L20)

`type Output struct { ... }`

Output is the result of comparing two composer.lock files.

#### func [ComposerDiff](/manager.go#L9)

`func ComposerDiff(input *Input) (*Output, error)`

#### func [Diff](/analyzer.go#L9)

`func Diff(previous, current shared.PackageMap) (*Output, error)`

Diff compares two composer.lock files and returns the differences
Note: Currently handles both lock file comparison AND requirement file integration.

### type [PackageChange](/types.go#L25)

`type PackageChange struct { ... }`

PackageChange contains detailed information about a package difference.

### type [PkgManagerInput](/types.go#L9)

`type PkgManagerInput struct { ... }`

## Sub Packages

* [composer](./composer)

* [shared](./shared)

* [shared_test](./shared_test)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
