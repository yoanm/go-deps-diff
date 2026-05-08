package semver_test

import (
	"fmt"
	"testing"

	"github.com/yoanm/go-deps-diff/contract/semver"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		version  string
		expected *semver.Version
	}{
		{name: "simple semver", version: "1.2.3", expected: &semver.Version{Major: 1, Minor: 2, Patch: 3, Extra: ""}},
		{name: "with v prefix", version: "v2.1.3", expected: &semver.Version{Major: 2, Minor: 1, Patch: 3, Extra: ""}},
		{name: "with prerelease", version: "1.2.3-beta.1", expected: &semver.Version{Major: 1, Minor: 2, Patch: 3, Extra: "-beta.1"}},
		{name: "with build metadata", version: "2.5.1+build.1", expected: &semver.Version{Major: 2, Minor: 5, Patch: 1, Extra: "+build.1"}},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			version, err := semver.Parse(testCase.version)
			if err != nil {
				t.Error(fmt.Errorf("an error occurred: %w", err))
			} else if err2 := difftesting.ValidateSemverVersion(version, testCase.expected); err2 != nil {
				t.Error(fmt.Errorf("checks failed: %w", err2))
			}
		})
	}
}

func TestParseSemverVersion_Error(t *testing.T) {
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

			_, err := semver.Parse(testCase.version)
			if err == nil {
				t.Errorf("an error is expected")
			} else if err2 := testCase.check(err); err2 != nil {
				t.Error(err2)
			}
		})
	}
}

func TestIsSemverValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		version string
	}{
		{
			name:    "simple semver",
			version: "1.2.3",
		},
		{
			name:    "with v prefix",
			version: "v2.1.3",
		},
		{
			name:    "with prerelease",
			version: "1.2.3-beta.1",
		},
		{
			name:    "with build metadata",
			version: "2.5.1+build.1",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			if !semver.IsValid(testCase.version) {
				t.Errorf("value is expected to be valid")
			}
		})
	}
}

func TestIsSemverValid_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		version string
	}{
		{
			name:    "invalid semver - commit hash",
			version: "abc123def456",
		},
		{
			name:    "invalid semver - dev-master",
			version: "dev-master",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			if semver.IsValid(testCase.version) {
				t.Errorf("value is expected to be invalid")
			}
		})
	}
}

func TestInvalidSemverComponentError(t *testing.T) {
	t.Parallel()

	err := semver.InvalidComponentError{Version: "dev-master"}

	if err.Error() != "invalid semver component: dev-master" {
		t.Error(fmt.Errorf("unexpected error: %w", err))
	}
}
