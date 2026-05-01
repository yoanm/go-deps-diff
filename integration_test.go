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
	}{
		{
			name:         "Simple",
			reqFilePath:  "./composer/testdata/composer-simple.json",
			lockFilePath: "./composer/testdata/composer-simple.lock",
		},
		{
			name:         "Complex",
			reqFilePath:  "./composer/testdata/composer-complex.json",
			lockFilePath: "./composer/testdata/composer-complex.lock",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			reqContent, err := os.ReadFile(testCase.reqFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading requirement file = %v", err)

				return
			}

			lockContent, err := os.ReadFile(testCase.lockFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading lock file = %v", err)

				return
			}

			out, err := depsdiff.ComposerDiff(&depsdiff.Input{
				Current: depsdiff.PkgManagerInput{
					Lock:        lockContent,
					Requirement: reqContent,
				},
				Previous: depsdiff.PkgManagerInput{
					Lock:        lockContent,
					Requirement: reqContent,
				},
			})
			if err != nil {
				t.Errorf("Diff() error = %v", err)

				return
			}

			if len(out) == 0 {
				t.Fatal("Diff() result check failed: no packages found in the output")
			}

			for _, change := range out {
				if change.Operation.Name != depsdiff.NoneOperation {
					t.Errorf(
						"Diff() result check failed: expected all packages to be unchanged, but %s isn't",
						change.Package.GetName(),
					)
				}
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

	if !reflect.DeepEqual(actualOperation, expectedOperation) {
		return fmt.Errorf("unexpected differences. Expected: %+v, Actual: %+v", expectedOperation, actualOperation)
	}

	return nil
}
