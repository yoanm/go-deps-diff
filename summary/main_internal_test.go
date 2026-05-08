package summary

import (
	"slices"
	"testing"

	"github.com/yoanm/go-deps-diff/contract"
	"github.com/yoanm/go-deps-diff/contract/semver"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func Test_buildItemMrkRowCells(t *testing.T) {
	t.Parallel()

	pkg := difftesting.GetDummyPackage()
	prevVersion := contract.PkgVersion{Raw: "1.2.3", Label: "1.2.3", Semver: &semver.Version{Major: 1, Minor: 2, Patch: 3, Extra: ""}}

	tests := []struct {
		name     string
		change   *contract.PackageChange
		mode     pkgRowMode
		expected []string
	}{
		{
			name: "Name, version, operation and previous version - operation with previous version",
			change: &contract.PackageChange{
				Package:         pkg,
				Operation:       difftesting.DowngradeMajorOp,
				PreviousVersion: prevVersion,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(prevVersion),
				buildOperationHTMLCell(difftesting.DowngradeMajorOp, 0),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Addition",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.AdditionOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildOperationHTMLCell(difftesting.AdditionOp, 2),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Removal",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.RemovalOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(difftesting.RemovalOp, 2),
			},
		},
		{
			name: "Name, version, operation and previous version - operation with only one version - Same",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.SameOp,
			},
			mode: fullPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(difftesting.SameOp, 2),
			},
		},
		{
			// Case is not actually expected to happen (it should be withOperationPkgRowMode mode in that case)
			name: "Name, version and operation - operation with previous version",
			change: &contract.PackageChange{
				Package:         pkg,
				Operation:       difftesting.DowngradeMajorOp,
				PreviousVersion: prevVersion,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(difftesting.DowngradeMajorOp, 0),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Addition",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.AdditionOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildOperationHTMLCell(difftesting.AdditionOp, 0),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Removal",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.RemovalOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(difftesting.RemovalOp, 0),
			},
		},
		{
			name: "Name, version and operation - operation with only one version - Same",
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.SameOp,
			},
			mode: withOperationPkgRowMode,
			expected: []string{
				// Rely on helper methods, goal here is to check the cell count and order, not the content !
				buildPackageNameHTMLCell(pkg),
				buildPackageVersionHTMLCell(pkg.GetVersion()),
				buildOperationHTMLCell(difftesting.SameOp, 0),
			},
		},
		{
			// Case is not actually expected to happen (it should be withOperationPkgRowMode mode in that case)
			name: "Name and version only - operation with previous version",
			change: &contract.PackageChange{
				Package:         pkg,
				Operation:       difftesting.DowngradeMajorOp,
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
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.AdditionOp,
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
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.RemovalOp,
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
			change: &contract.PackageChange{ // //nolint:exhaustruct // Useless for the test purpose
				Package:   pkg,
				Operation: difftesting.SameOp,
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
		&contract.PackageChange{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
		_testInvalidPkgRowMode,
	)

	t.Fatal("The code did not panic")
}
