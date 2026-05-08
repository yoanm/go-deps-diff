package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/contract"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func Test_guessShortestPkgRowMode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		pkgList  pkgList
		expected pkgRowMode
	}{
		{
			name: "Only Addition - Version + Operation",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Removal - Version + Operation",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Same - Version + Operation",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same - Version + Operation",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
			},
			expected: withOperationPkgRowMode,
		},

		{
			name: "Mix of Addition and Update - Full",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Removal and Update - Full",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Same and Update - Full",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same and Update - Version + Operation",
			pkgList: pkgList{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
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

func Test_buildSectionCounters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		categories subCategoriesMap
		expected   sectionSummaryCntMap
	}{
		{
			name:       "Base",
			categories: _testBuildSectionCountersBaseCategories,
			expected:   _testBuildSectionCountersBaseExpectations,
		},
		{
			name:       "SemverExtraUpdate icon if not basic Unknown",
			categories: _testBuildSectionCountersSemverExtraUpdateFallbackCategories,
			expected:   _testBuildSectionCountersSemverExtraUpdateFallbackExpectations,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := buildItemsCounters(testCase.categories)

			for key, val := range testCase.expected {
				currentVal, currentExists := current[key]
				switch {
				case !currentExists:
					t.Errorf("expected key %s not found", key)
				case val.count != currentVal.count:
					t.Errorf("unexpected count: got %d, want %d", val.count, val.count)
				case val.title != currentVal.title:
					t.Errorf("unexpected title: got %s, want %s", val.title, val.title)
				}
			}
		})
	}
}

var (
	//nolint:exhaustive // Meaningless in the test context
	_testBuildSectionCountersBaseCategories = map[markdownSubCategory]itemsMap{
		// Actual category  and content should not matter here !
		requirementSubCategory: map[markdownItem]pkgList{
			additionItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.AdditionOp},
			},
			removalItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.RemovalOp},
			},
			sameItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SameOp},
			},
			semverMajorUpgradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMajorOp},
			},
			semverMinorUpgradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradeMinorOp},
			},
			semverPatchUpgradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UpgradePatchOp},
			},
			semverMajorDowngradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMajorOp},
			},
			semverMinorDowngradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradeMinorOp},
			},
			semverPatchDowngradeItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.DowngradePatchOp},
			},
			unknownUpdateItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UnknownUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidDowngradeOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidUpgradeOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.UnknownUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidDowngradeOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.InvalidUpgradeOp},
			},
		},
	}
	_testBuildSectionCountersBaseExpectations = map[markdownItem]*itemsCounter{
		additionItem:             {title: "➕️", count: 1},
		removalItem:              {title: "❌", count: 2},
		sameItem:                 {title: "🟰", count: 3},
		semverMajorUpgradeItem:   {title: "<sub><sup>🔺.🔹.🔹</sup></sub>", count: 4},
		semverMinorUpgradeItem:   {title: "<sub><sup>🔹.🔺.🔹</sup></sub>", count: 5},
		semverPatchUpgradeItem:   {title: "<sub><sup>🔹.🔹.🔺</sup></sub>", count: 6},
		semverMajorDowngradeItem: {title: "<sub><sup>🔻.🔹.🔹</sup></sub>", count: 7},
		semverMinorDowngradeItem: {title: "<sub><sup>🔹.🔻.🔹</sup></sub>", count: 8},
		semverPatchDowngradeItem: {title: "<sub><sup>🔹.🔹.🔻</sup></sub>", count: 9},
		unknownUpdateItem:        {title: "❓", count: 10},
	}

	//nolint:exhaustive // Meaningless in the test context
	_testBuildSectionCountersSemverExtraUpdateFallbackCategories = map[markdownSubCategory]itemsMap{
		// Actual category  and content should not matter here !
		requirementSubCategory: map[markdownItem]pkgList{ //nolint:exhaustive // Meaningless in the test context
			unknownUpdateItem: []*contract.PackageChange{
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
				{Package: difftesting.GetDummyPackage(), Operation: difftesting.SemverExtraUpdateOp},
			},
		},
	}
	//nolint:exhaustive // Meaningless in the test context
	_testBuildSectionCountersSemverExtraUpdateFallbackExpectations = map[markdownItem]*itemsCounter{
		unknownUpdateItem: {title: "<sub><sup>🔹.🔹.🔹❓</sup></sub>", count: 5},
	}
)
