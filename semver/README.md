# semver

## Functions

### func [IsValid](./main.go#L26)

`func IsValid(value string) bool`

### func [Parse](./main.go#L32)

`func Parse(version string) (*contract.Semver, error)`

Parse parses a semantic Version string
Returns nil if parsing fails.

## Types

### type [InvalidComponentError](./main.go#L18)

`type InvalidComponentError struct { ... }`

#### func (InvalidComponentError) [Error](./main.go#L22)

`func (e InvalidComponentError) Error() string`

### type [InvalidVersionError](./main.go#L10)

`type InvalidVersionError struct { ... }`

#### func (InvalidVersionError) [Error](./main.go#L14)

`func (e InvalidVersionError) Error() string`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
