package semver

import (
	"regexp"
	"strconv"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Extra string
}

type InvalidVersionError struct {
	Version string
}

func (e InvalidVersionError) Error() string {
	return "invalid semver Version: " + e.Version
}

type InvalidComponentError struct {
	Version string
}

func (e InvalidComponentError) Error() string {
	return "invalid semver component: " + e.Version
}

func IsValid(value string) bool {
	return validateVersionRegexp.MatchString(value)
}

// Parse parses a semantic Version string
// Returns nil if parsing fails.
func Parse(version string) (*Version, error) {
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

	return &Version{Major: major, Minor: minor, Patch: patch, Extra: extra}, nil
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
