# shared

## Functions

### func [IsSemverValid](./semver.go#L31)

`func IsSemverValid(value string) bool`

## Types

### type [InvalidSemverComponentError](./semver.go#L23)

`type InvalidSemverComponentError struct { ... }`

#### func (InvalidSemverComponentError) [Error](./semver.go#L27)

`func (e InvalidSemverComponentError) Error() string`

### type [InvalidSemverVersionError](./semver.go#L15)

`type InvalidSemverVersionError struct { ... }`

#### func (InvalidSemverVersionError) [Error](./semver.go#L19)

`func (e InvalidSemverVersionError) Error() string`

### type [PackageMap](./wrapper_types.go#L5)

`type PackageMap map[string]PkgWrapper`

PackageMap holds package information for efficient lookup
Key is the package name (e.g., "vendor/package"), value is a wrapper providing package details and helper methods.

### type [PkgVersion](./wrapper_types.go#L28)

`type PkgVersion struct { ... }`

### type [PkgWrapper](./wrapper_types.go#L7)

`type PkgWrapper interface { ... }`

### type [SemverVersion](./semver.go#L8)

`type SemverVersion struct { ... }`

#### func [ParseSemverVersion](./semver.go#L38)

`func ParseSemverVersion(version string) (*SemverVersion, error)`

ParseSemverVersion parses a semantic Version string
Returns nil if parsing fails.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
