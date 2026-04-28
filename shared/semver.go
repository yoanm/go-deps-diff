package shared

import (
	"regexp"
	"strconv"
)

type SemverVersion struct {
	Major int
	Minor int
	Patch int
	Extra string
}

type InvalidSemverVersionError struct {
	Version string
}

func (e InvalidSemverVersionError) Error() string {
	return "invalid semver Version: " + e.Version
}

type InvalidSemverComponentError struct {
	Version string
}

func (e InvalidSemverComponentError) Error() string {
	return "invalid semver component: " + e.Version
}

func IsSemverValid(value string) bool {
	// Pattern: optional 'v', then MAJOR.MINOR.PATCH, then optional extra data (e.g. -beta, +build, etc.)
	return regexp.MustCompile(`^v?\d+\.\d+\.\d+.*$`).MatchString(value)
}

// ParseSemverVersion parses a semantic Version string
// Returns nil if parsing fails.
func ParseSemverVersion(version string) (*SemverVersion, error) {
	// Pattern: optional 'v', then MAJOR.MINOR.PATCH, then optional extra
	matches := regexp.MustCompile(`^(v)?(\d+)\.(\d+)\.(\d+)(.*)$`).FindStringSubmatch(version)
	if matches == nil {
		return nil, &InvalidSemverVersionError{Version: version}
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
		return nil, &InvalidSemverComponentError{Version: version}
	}

	return &SemverVersion{
		Major: major,
		Minor: minor,
		Patch: patch,
		Extra: extra,
	}, nil
}
