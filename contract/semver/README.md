# semver

## Functions

### func [IsValid](./main.go#L31)

`func IsValid(value string) bool`

## Types

### type [InvalidComponentError](./main.go#L23)

`type InvalidComponentError struct { ... }`

#### func (InvalidComponentError) [Error](./main.go#L27)

`func (e InvalidComponentError) Error() string`

### type [InvalidVersionError](./main.go#L15)

`type InvalidVersionError struct { ... }`

#### func (InvalidVersionError) [Error](./main.go#L19)

`func (e InvalidVersionError) Error() string`

### type [Version](./main.go#L8)

`type Version struct { ... }`

#### func [Parse](./main.go#L37)

`func Parse(version string) (*Version, error)`

Parse parses a semantic Version string
Returns nil if parsing fails.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
