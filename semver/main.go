// Package semver provides utilities for parsing and validating semantic versions.
//
// Semantic versioning follows the MAJOR.MINOR.PATCH[+extra] format (optionally prefixed with 'v').
// Examples: "1.2.3", "v1.2.3-beta", "2.0.0+build.123"
//
// This package is used internally by depsdiff to analyze version changes and provide
// semantic version type information (MAJOR, MINOR, PATCH, EXTRA updates).
package semver

import (
	"regexp"
	"strconv"

	"github.com/yoanm/go-deps-diff/contract"
)

// InvalidVersionError is returned when a version string does not match semantic version format.
type InvalidVersionError struct {
	Version string
}

func (e InvalidVersionError) Error() string {
	return "invalid semver Version: " + e.Version
}

// InvalidComponentError is returned when semantic version components cannot be parsed as integers.
type InvalidComponentError struct {
	Version string
}

func (e InvalidComponentError) Error() string {
	return "invalid semver component: " + e.Version
}

// IsValid reports whether value is a valid semantic version string.
//
// Valid formats include:
//   - "1.2.3"
//   - "v1.2.3" (with optional 'v' prefix)
//   - "1.2.3-beta" (with pre-release identifier)
//   - "1.2.3+build.123" (with build metadata)
//   - "1.2.3-beta+build.123" (with both)
//
// Returns true if the version matches semantic version format, false otherwise.
func IsValid(value string) bool {
	return validateVersionRegexp.MatchString(value)
}

// Parse parses a semantic version string and returns its components.
//
// The input string may optionally be prefixed with 'v' (e.g., "v1.2.3" or "1.2.3").
// The version string must follow semantic version format: MAJOR.MINOR.PATCH[+extra].
//
// Parameters:
//   - version: A semantic version string to parse
//
// Returns:
//   - *Semver: Pointer to parsed version components (Major, Minor, Patch, Extra), or nil on error
//   - error: InvalidVersionError if the version string format is invalid,
//     or InvalidComponentError if version components cannot be parsed as integers
//
// Example:
//
//	semver, err := Parse("1.2.3-beta")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Major: %d, Minor: %d\n", semver.Major, semver.Minor)
func Parse(version string) (*contract.Semver, error) {
	matches := parseVersionRegexp.FindStringSubmatch(version)
	if matches == nil {
		return nil, &InvalidVersionError{Version: version}
	}

	// matches[0] = full match
	// matches[1] = optional v
	// matches[2] = major
	// matches[3] = minor
	// matches[4] = patch
	// matches[5] = extra (may be empty)
	major, majorErr := strconv.Atoi(matches[2])
	minor, minorErr := strconv.Atoi(matches[3])
	patch, patchErr := strconv.Atoi(matches[4])
	extra := matches[5]

	if majorErr != nil || minorErr != nil || patchErr != nil {
		// Unlikely to happen as long as the regex is expecting digits, but we should still check for it
		return nil, &InvalidComponentError{Version: version}
	}

	return &contract.Semver{Major: major, Minor: minor, Patch: patch, Extra: extra}, nil
}

/**
	# SEMVER PATTERN
	- optional 'v'
    - MAJOR.MINOR.PATCH (all digits)
    - optional extra string (e.g. -beta, +build, etc.)
*/

// validateVersionRegexp will only validate the version.
// In order to extract component, see parseVersionRegexp instead.
var validateVersionRegexp = regexp.MustCompile(`^v?\d+\.\d+\.\d+.*$`)

// parseVersionRegexp will extract semver components (e.g. major, minor, path, extra) from the version.
// In order to simply validate a version, see validateVersionRegexp instead.
var parseVersionRegexp = regexp.MustCompile(`^(v)?(\d+)\.(\d+)\.(\d+)(.*)$`)
