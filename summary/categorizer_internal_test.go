package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/contract"

	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func Test_getMarkdownItemType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		change   *contract.PackageChange
		expected markdownItem
	}{
		{
			name: "UnknownUpdateOperation",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.UnknownUpdateOp,
			},
			expected: unknownUpdateItem,
		},
		{
			name: "RemovalOperation",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.RemovalOp,
			},
			expected: removalItem,
		},
		{
			name: "AdditionOperation",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.AdditionOp,
			},
			expected: additionItem,
		},
		{
			name: "NoChangeOperation",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.SameOp,
			},
			expected: sameItem,
		},
		{
			name: "Upgrade - SemverMajorUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.UpgradeMajorOp,
			},
			expected: semverMajorUpgradeItem,
		},
		{
			name: "Upgrade - SemverMinorUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.UpgradeMinorOp,
			},
			expected: semverMinorUpgradeItem,
		},
		{
			name: "Upgrade - SemverPatchUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.UpgradePatchOp,
			},
			expected: semverPatchUpgradeItem,
		},
		{
			name: "Upgrade - SemverExtraUpdate - not expected to exist, but should be categorized as unknown update",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverExtraUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Upgrade - SemverUnknownUpdate - not expected to exist, but should be categorized as unknown update",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: contract.Operation{Name: contract.UpgradeOperation, SemverType: contract.SemverUnknownUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Downgrade - SemverMajorUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.DowngradeMajorOp,
			},
			expected: semverMajorDowngradeItem,
		},
		{
			name: "Downgrade - SemverMinorUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.DowngradeMinorOp,
			},
			expected: semverMinorDowngradeItem,
		},

		{
			name: "Downgrade - SemverPatchUpdate",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.DowngradePatchOp,
			},
			expected: semverPatchDowngradeItem,
		},
		{
			name: "Downgrade - SemverExtraUpdate - not expected to exist, but should be categorized as unknown update",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverExtraUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Downgrade - SemverUnknownUpdate - not expected to exist, but should be categorized as unknown update",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: contract.Operation{Name: contract.DowngradeOperation, SemverType: contract.SemverUnknownUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Unknown for unmanaged case",
			change: &contract.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: difftesting.InvalidOp,
			},
			expected: unknownUpdateItem,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := getMarkdownItemType(testCase.change)

			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}
