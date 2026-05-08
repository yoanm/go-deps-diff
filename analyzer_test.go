package depsdiff_test

import (
	"fmt"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/contract"
	"github.com/yoanm/go-deps-diff/contract/semver"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func TestDiff_NoChange(t *testing.T) {
	t.Parallel()

	previous := map[string]contract.PkgWrapper{
		"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
			Name:    "vendor/pkg",
			Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		},
	}
	current := map[string]contract.PkgWrapper{
		"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
			Name:    "vendor/pkg",
			Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
		},
	}

	out, err := depsdiff.Diff(previous, current)
	if err != nil {
		t.Fatal(fmt.Errorf("error during diff process: %w", err))
	}

	change, pkgExists := out["vendor/pkg"]
	switch {
	case !pkgExists:
		t.Fatal("package 'vendor/pkg' is expected in the package map")
	case change.Operation.Name != contract.NoChangeOperation:
		t.Fatalf(
			"unexpected Operation: got %s, want %s",
			change.Operation.Name,
			contract.NoChangeOperation,
		)
	case change.Operation.SemverType != contract.SemverNoUpdate:
		t.Fatalf(
			"unexpected SemverType: got %s, want %s",
			change.Operation.SemverType,
			contract.SemverNoUpdate,
		)
	}
}

func TestDiff_BasicComparison(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		previous              contract.PackageMap
		current               contract.PackageMap
		expectedOperationName contract.OperationName
		expectedSemverType    contract.OperationSemverType
	}{
		{
			name:     "added package",
			previous: map[string]contract.PkgWrapper{},
			current: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			expectedOperationName: contract.AdditionOperation,
			expectedSemverType:    contract.SemverNoUpdate,
		},
		{
			name: "unchanged package",
			previous: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			current: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			expectedOperationName: contract.NoChangeOperation,
			expectedSemverType:    contract.SemverNoUpdate,
		},
		{
			name: "removed package",
			previous: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			current:               map[string]contract.PkgWrapper{},
			expectedOperationName: contract.RemovalOperation,
			expectedSemverType:    contract.SemverNoUpdate,
		},
		{
			name: "upgraded package",
			previous: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			current: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "2.0.0", Label: "2.0.0", Semver: &semver.Version{Major: 2, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			expectedOperationName: contract.UpgradeOperation,
			expectedSemverType:    contract.SemverMajorUpdate,
		},
		{
			name: "downgraded package",
			previous: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.1.0", Label: "1.1.0", Semver: &semver.Version{Major: 1, Minor: 1, Patch: 0, Extra: ""}},
				},
			},
			current: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
				},
			},
			expectedOperationName: contract.DowngradeOperation,
			expectedSemverType:    contract.SemverMinorUpdate,
		},
		{
			name: "unknown update",
			previous: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "abcdef", Label: "dev-master#abcdef", Semver: nil},
				},
			},
			current: map[string]contract.PkgWrapper{
				"vendor/pkg": &difftesting.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: contract.PkgVersion{Raw: "fedcba", Label: "1.0.0#fedcba", Semver: nil},
				},
			},

			expectedOperationName: contract.UnknownUpdateOperation,
			expectedSemverType:    contract.SemverUnknownUpdate,
		},
	}

	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			out, err := depsdiff.Diff(testData.previous, testData.current)
			if err != nil {
				t.Fatal(fmt.Errorf("error during diff process: %w", err))
			}

			change, pkgExists := out["vendor/pkg"]
			switch {
			case !pkgExists:
				t.Fatal("package 'vendor/pkg' is expected in the package map")
			case change.Operation.Name != testData.expectedOperationName:
				t.Fatalf(
					"unexpected Operation: got %s, want %s",
					change.Operation.Name,
					testData.expectedOperationName,
				)
			case change.Operation.SemverType != testData.expectedSemverType:
				t.Fatalf(
					"unexpected SemverType: got %s, want %s",
					change.Operation.SemverType,
					testData.expectedSemverType,
				)
			}
		})
	}
}
