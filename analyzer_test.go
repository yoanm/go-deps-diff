package depsdiff_test

import (
	"errors"
	"fmt"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func TestDiff_NoChange(t *testing.T) {
	t.Parallel()

	previous := map[string]shared.PkgWrapper{
		"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
			Name:    "vendor/pkg",
			Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
		},
	}
	current := map[string]shared.PkgWrapper{
		"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
			Name:    "vendor/pkg",
			Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
		},
	}

	out, err := depsdiff.Diff(previous, current)
	if err != nil {
		t.Fatal(fmt.Errorf("error during diff process: %w", err))
	}

	if len(out) != 0 {
		t.Fatal(errors.New("expected no changes, but got some"))
	}
}

func TestDiff_BasicComparison(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		previous              shared.PackageMap
		current               shared.PackageMap
		expectedOperationName depsdiff.OperationName
		expectedSemverType    depsdiff.OperationSemverType
	}{
		{
			name:     "added package",
			previous: map[string]shared.PkgWrapper{},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
				},
			},
			expectedOperationName: depsdiff.AdditionOperation,
			expectedSemverType:    depsdiff.SemverNoUpdate,
		},
		{
			name: "removed package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
				},
			},
			current:               map[string]shared.PkgWrapper{},
			expectedOperationName: depsdiff.RemovalOperation,
			expectedSemverType:    depsdiff.SemverNoUpdate,
		},
		{
			name: "upgraded package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "2.0.0", Label: "2.0.0"},
				},
			},
			expectedOperationName: depsdiff.UpgradeOperation,
			expectedSemverType:    depsdiff.SemverMajorUpdate,
		},
		{
			name: "downgraded package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.1.0", Label: "1.1.0"},
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0"},
				},
			},
			expectedOperationName: depsdiff.DowngradeOperation,
			expectedSemverType:    depsdiff.SemverMinorUpdate,
		},
		{
			name: "unknown update",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "abcdef", Label: "dev-master#abcdef"},
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "fedcba", Label: "1.0.0#fedcba"},
				},
			},

			expectedOperationName: depsdiff.UnknownUpdateOperation,
			expectedSemverType:    depsdiff.SemverUnknownUpdate,
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
