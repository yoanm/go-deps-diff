package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
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
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Removal - Version + Operation",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Only Same - Version + Operation",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
			},
			expected: withOperationPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same - Version + Operation",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
			},
			expected: withOperationPkgRowMode,
		},

		{
			name: "Mix of Addition and Update - Full",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Removal and Update - Full",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Same and Update - Full",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
			},
			expected: fullPkgRowMode,
		},
		{
			name: "Mix of Addition/Removal/Same and Update - Version + Operation",
			pkgList: pkgList{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
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

//nolint:gochecknoglobals // Just to keep it outside the function
var (
	//nolint:exhaustive // Meaningless in the test context
	_testBuildSectionCountersBaseCategories = map[markdownSubCategory]itemsMap{
		// Actual category  and content should not matter here !
		requirementSubCategory: map[markdownItem]pkgList{
			additionItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.AdditionOp},
			},
			removalItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.RemovalOp},
			},
			sameItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SameOp},
			},
			semverMajorUpgradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMajorOp},
			},
			semverMinorUpgradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradeMinorOp},
			},
			semverPatchUpgradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UpgradePatchOp},
			},
			semverMajorDowngradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMajorOp},
			},
			semverMinorDowngradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradeMinorOp},
			},
			semverPatchDowngradeItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.DowngradePatchOp},
			},
			unknownUpdateItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UnknownUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidDowngradeOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidUpgradeOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.UnknownUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidDowngradeOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.InvalidUpgradeOp},
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
			unknownUpdateItem: []*shared.PackageChange{
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
				{Package: shared_test.GetDummyPackage(), Operation: shared_test.SemverExtraUpdateOp},
			},
		},
	}
	//nolint:exhaustive // Meaningless in the test context
	_testBuildSectionCountersSemverExtraUpdateFallbackExpectations = map[markdownItem]*itemsCounter{
		unknownUpdateItem: {title: "<sub><sup>🔹.🔹.🔹❓</sup></sub>", count: 5},
	}
)
