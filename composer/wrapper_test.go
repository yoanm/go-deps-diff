package composer_test

import (
	"fmt"
	"testing"

	"github.com/yoanm/go-deps-diff/composer"
)

func TestBuildMapFromBytes_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		reqData  []byte
		lockData []byte
		checkFn  func(err error) error
	}{
		{
			name:    "invalid json - req file",
			reqData: []byte(`{invalid}`),
			lockData: []byte(`{
				"packages": [
					{"name": "vendor/pkg", "version": "1.0.0"}
				]
			}`),
			checkFn: func(err error) error {
				//nolint:lll // Doesn't make any sense to refactor this just to avoid a long line in a test case
				if err.Error() != "parsing requirement file content: invalid JSON: invalid character 'i' looking for beginning of object key string" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "invalid json - lock file",
			reqData: []byte(`{
				"require": {"vendor/pkg": "^1.0"},
				"require-dev": {"vendor/test": "^1.0"}
			}`),
			lockData: []byte(`{invalid}`),
			checkFn: func(err error) error {
				//nolint:lll // Doesn't make any sense to refactor this just to avoid a long line in a test case
				if err.Error() != "parsing lock file content: invalid JSON: invalid character 'i' looking for beginning of object key string" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name:    "empty input - req file",
			reqData: []byte{},
			lockData: []byte(`{
				"packages": [
					{"name": "vendor/pkg", "version": "1.0.0"}
				]
			}`),
			checkFn: func(err error) error {
				if err.Error() != "parsing requirement file content: invalid format: empty input" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "empty input - lock file",
			reqData: []byte(`{
				"require": {"vendor/pkg": "^1.0"},
				"require-dev": {"vendor/test": "^1.0"}
			}`),
			lockData: []byte{},
			checkFn: func(err error) error {
				if err.Error() != "parsing lock file content: invalid format: empty input" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name:    "missing require arrays - req file",
			reqData: []byte(`{"other": "field"}`),
			lockData: []byte(`{
				"packages": [
					{"name": "vendor/pkg", "version": "1.0.0"}
				]
			}`),
			checkFn: func(err error) error {
				if err.Error() != "parsing requirement file content: invalid format: missing 'require' or 'require-dev' fields" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "missing require arrays - lock file",
			reqData: []byte(`{
				"require": {"vendor/pkg": "^1.0"},
				"require-dev": {"vendor/test": "^1.0"}
			}`),
			lockData: []byte(`{"other": "field"}`),
			checkFn: func(err error) error {
				if err.Error() != "parsing lock file content: invalid format: missing 'packages' or 'packages-dev' fields" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, err := composer.BuildMapFromBytes(testCase.reqData, testCase.lockData)
			if err == nil {
				t.Errorf("an error is expected")
			} else if err2 := testCase.checkFn(err); err2 != nil {
				t.Error(err2)
			}
		})
	}
}
