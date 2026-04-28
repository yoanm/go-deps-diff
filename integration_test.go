package depsdiff_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func TestIntegration_Fixtures(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		reqFilePath  string
		lockFilePath string
		checkFn      func(out *depsdiff.Output) bool
	}{
		{
			name:         "Simple",
			reqFilePath:  "./composer/testdata/composer-simple.json",
			lockFilePath: "./composer/testdata/composer-simple.lock",
			checkFn: func(out *depsdiff.Output) bool {
				return len(out.Changes) == 0
			},
		},
		{
			name:         "Complex",
			reqFilePath:  "./composer/testdata/composer-complex.json",
			lockFilePath: "./composer/testdata/composer-complex.lock",
			checkFn: func(out *depsdiff.Output) bool {
				return len(out.Changes) == 0
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// Load fixture files
			simpleReq, err := os.ReadFile(testCase.reqFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading requirement file = %v", err)

				return
			}

			simpleLock, err := os.ReadFile(testCase.lockFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading lock file = %v", err)

				return
			}

			out, err := depsdiff.ComposerDiff(&depsdiff.Input{
				Current: depsdiff.PkgManagerInput{
					Lock:        simpleLock,
					Requirement: simpleReq,
				},
				Previous: depsdiff.PkgManagerInput{
					Lock:        simpleLock,
					Requirement: simpleReq,
				},
			})
			if err != nil {
				t.Errorf("Diff() error = %v", err)

				return
			}

			if testCase.checkFn != nil && !testCase.checkFn(out) {
				t.Errorf("Diff() result check failed")
			}
		})
	}
}

func validateChanges(actual, expectedChanges map[string]depsdiff.PackageChange) []error {
	var errList []error

	for _, expectedChange := range expectedChanges {
		pkgName := expectedChange.Package.GetName()
		actualChange, exists := actual[pkgName]

		if !exists {
			errList = append(errList, fmt.Errorf("Package %s is expected to exist", pkgName))
		} else {
			if err := shared_test.ValidateWrapperPackage(actualChange.Package, expectedChange.Package); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected Package differences: %w", pkgName, err))
			}

			if err := ValidateWrapperOperation(actualChange.Operation, expectedChange.Operation); err != nil {
				errList = append(errList, fmt.Errorf("package %s has unexpected Operation differences: %w", pkgName, err))
			}

			if err := shared_test.ValidatePackageVersion(actualChange.PreviousVersion, expectedChange.PreviousVersion); err != nil { //nolint:lll // Meaningless here
				errList = append(errList, fmt.Errorf("package %s has unexpected PreviousVersion differences: %w", pkgName, err))
			}
		}
	}

	for pkgName := range actual {
		if change, exists := expectedChanges[pkgName]; !exists {
			errList = append(errList, fmt.Errorf("package %s is not expected to exist. %+v", pkgName, change))
		}
	}

	return errList
}

func ValidateWrapperOperation(actualOperation, expectedOperation depsdiff.Operation) error {
	if actualOperation.Name != expectedOperation.Name {
		return fmt.Errorf("unexpected Name value. Expected: %s Actual: %s", expectedOperation.Name, actualOperation.Name)
	}

	if actualOperation.SemverType != expectedOperation.SemverType {
		return fmt.Errorf(
			"unexpected SemverType value. Expected: %s Actual: %s",
			expectedOperation.SemverType,
			actualOperation.SemverType,
		)
	}

	if actualOperation.Direction != expectedOperation.Direction {
		return fmt.Errorf("unexpected Direction value. Expected: %s Actual: %s",
			expectedOperation.Direction,
			actualOperation.Direction,
		)
	}

	if !reflect.DeepEqual(actualOperation, expectedOperation) {
		return fmt.Errorf("unexpected differences. Expected: %+v, Actual: %+v", expectedOperation, actualOperation)
	}

	return nil
}
