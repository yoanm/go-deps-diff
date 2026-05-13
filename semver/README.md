# semver

Package semver provides utilities for parsing and validating semantic versions.

Semantic versioning follows the MAJOR.MINOR.PATCH[+extra] format (optionally prefixed with 'v').
Examples: "1.2.3", "v1.2.3-beta", "2.0.0+build.123"

This package is used internally by depsdiff to analyze version changes and provide
semantic version type information (MAJOR, MINOR, PATCH, EXTRA updates).

## Functions

### func [IsValid](./main.go#L45)

`func IsValid(value string) bool`

IsValid reports whether value is a valid semantic version string.

Valid formats include:

```diff
- "1.2.3"
- "v1.2.3" (with optional 'v' prefix)
- "1.2.3-beta" (with pre-release identifier)
- "1.2.3+build.123" (with build metadata)
- "1.2.3-beta+build.123" (with both)
```

Returns true if the version matches semantic version format, false otherwise.

### func [Parse](./main.go#L69)

`func Parse(version string) (*contract.Semver, error)`

Parse parses a semantic version string and returns its components.

The input string may optionally be prefixed with 'v' (e.g., "v1.2.3" or "1.2.3").
The version string must follow semantic version format: MAJOR.MINOR.PATCH[+extra].

Parameters:

```go
- version: A semantic version string to parse
```

Returns:

```diff
- *Semver: Pointer to parsed version components (Major, Minor, Patch, Extra), or nil on error
- error: InvalidVersionError if the version string format is invalid,
  or InvalidComponentError if version components cannot be parsed as integers
```

Example:

```go
semver, err := Parse("1.2.3-beta")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Major: %d, Minor: %d\n", semver.Major, semver.Minor)
```

## Types

### type [InvalidComponentError](./main.go#L27)

`type InvalidComponentError struct { ... }`

InvalidComponentError is returned when semantic version components cannot be parsed as integers.

#### func (InvalidComponentError) [Error](./main.go#L31)

`func (e InvalidComponentError) Error() string`

### type [InvalidVersionError](./main.go#L18)

`type InvalidVersionError struct { ... }`

InvalidVersionError is returned when a version string does not match semantic version format.

#### func (InvalidVersionError) [Error](./main.go#L22)

`func (e InvalidVersionError) Error() string`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
