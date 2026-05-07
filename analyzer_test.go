package depsdiff_test

import (
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
			Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
		},
	}
	current := map[string]shared.PkgWrapper{
		"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
			Name:    "vendor/pkg",
			Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
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
	case change.Operation.Name != shared.NoChangeOperation:
		t.Fatalf(
			"unexpected Operation: got %s, want %s",
			change.Operation.Name,
			shared.NoChangeOperation,
		)
	case change.Operation.SemverType != shared.SemverNoUpdate:
		t.Fatalf(
			"unexpected SemverType: got %s, want %s",
			change.Operation.SemverType,
			shared.SemverNoUpdate,
		)
	}
}

func TestDiff_BasicComparison(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		previous              shared.PackageMap
		current               shared.PackageMap
		expectedOperationName shared.OperationName
		expectedSemverType    shared.OperationSemverType
	}{
		{
			name:     "added package",
			previous: map[string]shared.PkgWrapper{},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			expectedOperationName: shared.AdditionOperation,
			expectedSemverType:    shared.SemverNoUpdate,
		},
		{
			name: "unchanged package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			expectedOperationName: shared.NoChangeOperation,
			expectedSemverType:    shared.SemverNoUpdate,
		},
		{
			name: "removed package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			current:               map[string]shared.PkgWrapper{},
			expectedOperationName: shared.RemovalOperation,
			expectedSemverType:    shared.SemverNoUpdate,
		},
		{
			name: "upgraded package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "2.0.0", Label: "2.0.0", Semver: &shared.SemverVersion{Major: 2, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			expectedOperationName: shared.UpgradeOperation,
			expectedSemverType:    shared.SemverMajorUpdate,
		},
		{
			name: "downgraded package",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.1.0", Label: "1.1.0", Semver: &shared.SemverVersion{Major: 1, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &shared.SemverVersion{Major: 1, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
				},
			},
			expectedOperationName: shared.DowngradeOperation,
			expectedSemverType:    shared.SemverMinorUpdate,
		},
		{
			name: "unknown update",
			previous: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "abcdef", Label: "dev-master#abcdef", Semver: nil},
				},
			},
			current: map[string]shared.PkgWrapper{
				"vendor/pkg": &shared_test.TestPkgWrapper{ //nolint:exhaustruct // Useless for the test purpose
					Name:    "vendor/pkg",
					Version: shared.PkgVersion{Raw: "fedcba", Label: "1.0.0#fedcba", Semver: nil},
				},
			},

			expectedOperationName: shared.UnknownUpdateOperation,
			expectedSemverType:    shared.SemverUnknownUpdate,
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
