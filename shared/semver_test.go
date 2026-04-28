package shared_test

import (
	"fmt"
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
)

func TestParseSemver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		version string
		check   func(version *shared.SemverVersion) error
	}{
		{
			name:    "simple semver",
			version: "1.2.3",
			check: func(version *shared.SemverVersion) error {
				return validateSemverVersion(version, 1, 2, 3, "")
			},
		},
		{
			name:    "with v prefix",
			version: "v2.1.3",
			check: func(version *shared.SemverVersion) error {
				return validateSemverVersion(version, 2, 1, 3, "")
			},
		},
		{
			name:    "with prerelease",
			version: "1.2.3-beta.1",
			check: func(version *shared.SemverVersion) error {
				return validateSemverVersion(version, 1, 2, 3, "-beta.1")
			},
		},
		{
			name:    "with build metadata",
			version: "2.5.1+build.1",
			check: func(version *shared.SemverVersion) error {
				return validateSemverVersion(version, 2, 5, 1, "+build.1")
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			tag, err := shared.ParseSemverVersion(testCase.version)
			if err != nil {
				t.Error(fmt.Errorf("an error occurred: %w", err))
			} else if err2 := testCase.check(tag); err2 != nil {
				t.Error(fmt.Errorf("checks failed: %w", err2))
			}
		})
	}
}

func validateSemverVersion(
	version *shared.SemverVersion,
	expectedMajor, expectedMinor,
	expectedPatch int,
	expectedExtra string,
) error {
	if version.Major != expectedMajor {
		return fmt.Errorf("unexpected major Version: got %d, want %d", version.Major, expectedMajor)
	}

	if version.Minor != expectedMinor {
		return fmt.Errorf("unexpected minor Version: got %d, want %d", version.Minor, expectedMinor)
	}

	if version.Patch != expectedPatch {
		return fmt.Errorf("unexpected patch Version: got %d, want %d", version.Patch, expectedPatch)
	}

	if version.Extra != expectedExtra {
		return fmt.Errorf("unexpected extra part: got '%s', want '%s'", version.Extra, expectedExtra)
	}

	return nil
}

func TestParseSemver_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		version string
		check   func(err error) error
	}{
		{
			name:    "invalid semver - commit hash",
			version: "abc123def456",
			check: func(err error) error {
				if err.Error() != "invalid semver Version: abc123def456" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name:    "invalid semver - dev-master",
			version: "dev-master",
			check: func(err error) error {
				if err.Error() != "invalid semver Version: dev-master" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, err := shared.ParseSemverVersion(testCase.version)
			if err == nil {
				t.Errorf("an error is expected")
			} else if err2 := testCase.check(err); err2 != nil {
				t.Error(err2)
			}
		})
	}
}
