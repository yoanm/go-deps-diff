package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func Test_guessShortestPkgRowMode(t *testing.T) {
	t.Parallel()

	dummyPackage := &shared_test.TestPkgWrapper{
		Name:               "vendor/package",
		Abandoned:          true,
		Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
		Link:               "",
		DevOnly:            false,
		RootRequirement:    false,
		RootDevRequirement: false,
	}
	additionOp := shared.Operation{Name: shared.AdditionOperation, SemverType: shared.SemverNoUpdate}
	removalOp := shared.Operation{Name: shared.RemovalOperation, SemverType: shared.SemverNoUpdate}
	sameOp := shared.Operation{Name: shared.NoChangeOperation, SemverType: shared.SemverNoUpdate}
	upgradeOp := shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMajorUpdate}

	tests := []struct {
		name     string
		pkgList  pkgList
		expected pkgRowMode
	}{
		{
			name: "Only Addition - Version + Operation",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: additionOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Removal - Version + Operation",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: removalOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Same - Version + Operation",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: sameOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same - Version + Operation",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: additionOp},
				{Package: dummyPackage, Operation: removalOp},
				{Package: dummyPackage, Operation: sameOp},
			},
			expected: withOperationPkgRowMode,
		},

		{
			name: "Mix of Addition and Update - Full",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: additionOp},
				{Package: dummyPackage, Operation: upgradeOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Removal and Update - Full",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: removalOp},
				{Package: dummyPackage, Operation: upgradeOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Same and Update - Full",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: sameOp},
				{Package: dummyPackage, Operation: upgradeOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same and Update - Version + Operation",
			pkgList: pkgList{
				{Package: dummyPackage, Operation: additionOp},
				{Package: dummyPackage, Operation: removalOp},
				{Package: dummyPackage, Operation: upgradeOp},
				{Package: dummyPackage, Operation: sameOp},
			},
			expected: fullPkgRowMode,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := guessShortestPkgRowMode(testCase.pkgList)

			if testCase.expected != current {
				t.Errorf("unexpected output: got %d, want %d", current, testCase.expected)
			}
		})
	}
}
