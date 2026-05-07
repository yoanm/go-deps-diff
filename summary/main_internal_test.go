package summary

import (
	"slices"
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func Test_buildItemMrkRowCells(t *testing.T) {
	t.Parallel()

	pkg := shared_test.GetDummyPackage()
	prevVersion := shared.PkgVersion{Raw: "1.2.3", Label: "1.2.3", Semver: &shared.SemverVersion{Major: 1, Minor: 2, Patch: 3, Extra: ""}} //nolint:lll // Meaningless for tests !

	tests := []struct {
		name     string
		change   *shared.PackageChange
		mode     pkgRowMode
		expected []string
	}{
		{
			name: "Name, version, operation and previous version - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: prevVersion,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(prevVersion),
				buildOperationHTMLCell(shared_test.DowngradeMajorOp, 0),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Addition",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildOperationHTMLCell(shared_test.AdditionOp, 2),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Removal",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.RemovalOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(shared_test.RemovalOp, 2),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Same",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.SameOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(shared_test.SameOp, 2),
			},
		},
		{
			// Case is not actually expected to happen (it should be withOperationPkgRowMode mode in that case)
			name: "Name, version and operation - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: prevVersion,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(shared_test.DowngradeMajorOp, 0),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Addition",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildOperationHTMLCell(shared_test.AdditionOp, 0),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Removal",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.RemovalOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(shared_test.RemovalOp, 0),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Same",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.SameOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(shared_test.SameOp, 0),
			},
		},
		{
			// Case is not actually expected to happen (it should be withOperationPkgRowMode mode in that case)
			name: "Name and version only - operation with previous version",
			change: &shared.PackageChange{
				Package:         pkg,
				Operation:       shared_test.DowngradeMajorOp,
				PreviousVersion: prevVersion,
			},
			mode: versionOnlyPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name and version only - operation with only one version - Addition",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.AdditionOp,
			},
			mode: versionOnlyPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name and version only - operation with only one version - Removal",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.RemovalOp,
			},
			mode: versionOnlyPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name and version only - operation with only one version - Same",
			change: &shared.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: shared_test.SameOp,
			},
			mode: versionOnlyPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := buildItemMrkRowCells(testCase.change, testCase.mode)

			if !slices.Equal(current, testCase.expected) {
				t.Errorf("unexpected result: got %s, want %s", current, testCase.expected)
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
