package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func Test_getMarkdownItemType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		change   *shared.PackageChange
		expected markdownItem
	}{
		{
			name: "UnknownUpdateOperation",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UnknownUpdateOperation, SemverType: shared.SemverUnknownUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "RemovalOperation",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.RemovalOperation, SemverType: shared.SemverNoUpdate},
			},
			expected: removalItem,
		},
		{
			name: "AdditionOperation",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.AdditionOperation, SemverType: shared.SemverNoUpdate},
			},
			expected: additionItem,
		},
		{
			name: "NoChangeOperation",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.NoChangeOperation, SemverType: shared.SemverNoUpdate},
			},
			expected: sameItem,
		},
		{
			name: "Upgrade - SemverMajorUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMajorUpdate},
			},
			expected: semverMajorUpgradeItem,
		},
		{
			name: "Upgrade - SemverMinorUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMinorUpdate},
			},
			expected: semverMinorUpgradeItem,
		},
		{
			name: "Upgrade - SemverPatchUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverPatchUpdate},
			},
			expected: semverPatchUpgradeItem,
		},
		{
			name: "Upgrade - SemverExtraUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverExtraUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Upgrade - SemverExtraUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverUnknownUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Downgrade - SemverMajorUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMajorUpdate},
			},
			expected: semverMajorDowngradeItem,
		},
		{
			name: "Downgrade - SemverMinorUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMinorUpdate},
			},
			expected: semverMinorDowngradeItem,
		},

		{
			name: "Downgrade - SemverPatchUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverPatchUpdate},
			},
			expected: semverPatchDowngradeItem,
		},
		{
			name: "Downgrade - SemverExtraUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverExtraUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Downgrade - SemverExtraUpdate",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				Operation: shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverUnknownUpdate},
			},
			expected: unknownUpdateItem,
		},
		{
			name: "Unknown for unmanaged case",
			change: &shared.PackageChange{ //nolint:exhaustruct // Useless for the test purpose
				// Following operation (downgrade + semver no update)is not expected to exist
				Operation: shared.Operation{Name: shared_test.InvalidOperationName, SemverType: shared.SemverNoUpdate},
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
