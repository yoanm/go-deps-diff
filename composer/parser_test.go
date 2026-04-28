package composer_test

import (
	"fmt"
	"testing"

	"github.com/yoanm/go-deps-diff/composer"
)

func TestParseLock(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		checkFn func(lock *composer.ComposerLock) bool
	}{
		{
			name: "valid simple lock",
			data: []byte(`{
				"packages": [
					{"name": "vendor/pkg", "version": "1.0.0"}
				]
			}`),
			checkFn: func(lock *composer.ComposerLock) bool {
				return lock != nil && len(lock.Packages) == 1 &&
					lock.Packages[0].Name == "vendor/pkg"
			},
		},
		{
			name: "valid lock with packages-dev",
			data: []byte(`{
				"packages": [],
				"packages-dev": [
					{"name": "vendor/test", "version": "1.0.0"}
				]
			}`),
			checkFn: func(lock *composer.ComposerLock) bool {
				return lock != nil && len(lock.PackagesDev) == 1 &&
					lock.PackagesDev[0].Name == "vendor/test"
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			lock, err := composer.ParseLock(testCase.data)
			if err != nil {
				t.Errorf("ParseLock() error = %v", err)
			}

			if !testCase.checkFn(lock) {
				t.Errorf("ParseLock() result check failed")
			}
		})
	}
}

func TestParseLock_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		checkFn func(err error) error
	}{
		{
			name: "invalid json",
			data: []byte(`{invalid}`),
			checkFn: func(err error) error {
				if err.Error() != "invalid JSON: invalid character 'i' looking for beginning of object key string" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "empty input",
			data: []byte{},
			checkFn: func(err error) error {
				if err.Error() != "invalid format: empty input" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "missing packages arrays",
			data: []byte(`{"other": "field"}`),
			checkFn: func(err error) error {
				if err.Error() != "invalid format: missing 'packages' or 'packages-dev' fields" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, err := composer.ParseLock(testCase.data)
			if err == nil {
				t.Errorf("an error is expected")
			} else if err2 := testCase.checkFn(err); err2 != nil {
				t.Error(err2)
			}
		})
	}
}

func TestParseReq(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		checkFn func(req *composer.ComposerReq) bool
	}{
		{
			name: "valid composer.json",
			data: []byte(`{
				"require": {"vendor/pkg": "^1.0"},
				"require-dev": {"vendor/test": "^1.0"}
			}`),
			checkFn: func(req *composer.ComposerReq) bool {
				return req != nil && req.Require["vendor/pkg"] == "^1.0" &&
					req.RequireDev["vendor/test"] == "^1.0"
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			lock, err := composer.ParseReq(testCase.data)
			if err != nil {
				t.Errorf("ParseReq() error = %v", err)
			}

			if !testCase.checkFn(lock) {
				t.Errorf("ParseReq() result check failed")
			}
		})
	}
}

func TestParseReq_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		checkFn func(err error) error
	}{
		{
			name: "invalid json",
			data: []byte(`{invalid}`),
			checkFn: func(err error) error {
				if err.Error() != "invalid JSON: invalid character 'i' looking for beginning of object key string" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "empty input",
			data: []byte{},
			checkFn: func(err error) error {
				if err.Error() != "invalid format: empty input" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
		{
			name: "missing require arrays",
			data: []byte(`{"other": "field"}`),
			checkFn: func(err error) error {
				if err.Error() != "invalid format: missing 'packages' or 'packages-dev' fields" {
					return fmt.Errorf("unexpected error: %w", err)
				}

				return nil
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, err := composer.ParseLock(testCase.data)
			if err == nil {
				t.Errorf("an error is expected")
			} else if err2 := testCase.checkFn(err); err2 != nil {
				t.Error(err2)
			}
		})
	}
}
