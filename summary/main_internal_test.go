package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func Test_buildItemMrkRowCells_count(t *testing.T) {
	t.Parallel()

	pkg := shared_test.GetDummyPackage()

	tests := []struct {
		name          string
		change        *shared.PackageChange
		mode          pkgRowMode
		expectedCount int
	}{
		{
			name: "Name, version, operation and previous version - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: shared.PkgVersion{Raw: "1.2.3", Label: "1.2.3"},
			},
			mode:          fullPkgRowMode,
			expectedCount: 4,
		},
		{
			name: "Name, version, operation and previous version - operation with only one version",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode:          fullPkgRowMode,
			expectedCount: 3,
		},
		{
			name: "Name, version and operation - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: shared.PkgVersion{Raw: "1.2.3", Label: "1.2.3"},
			},
			mode:          withOperationPkgRowMode,
			expectedCount: 3,
		},
		{
			name: "Name, version and operation - operation with only one version",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode:          withOperationPkgRowMode,
			expectedCount: 3,
		},
		{
			// Case is not actually expected to happen (it should be withOperationPkgRowMode mode in that case)
			name: "Name and version only - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: shared.PkgVersion{Raw: "1.2.3", Label: "1.2.3"},
			},
			mode:          versionOnlyPkgRowMode,
			expectedCount: 2,
		},
		{
			name: "Name and version only - operation with only one version",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode:          versionOnlyPkgRowMode,
			expectedCount: 2,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := buildItemMrkRowCells(testCase.change, testCase.mode)

			if len(current) != testCase.expectedCount {
				t.Errorf("unexpected cell count: got %d, want %d (cells:%s)", len(current), testCase.expectedCount, current)
			}
		})
	}
}

const _testInvalidPkgRowMode pkgRowMode = -1

func Test_buildItemMrkRowCells_panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Log("function thrown a panic as expected.")
		}
	}()

	_ = buildItemMrkRowCells(
		//nolint:exhaustruct // Useless for the test purpose
		&shared.PackageChange{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
		_testInvalidPkgRowMode,
	)

	t.Fatal("The code did not panic")
}
