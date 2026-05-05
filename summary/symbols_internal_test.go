package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
)

const _testUnmanagedOperation shared.OperationName = "ARGH"

func Test_getOperationSymbol(t *testing.T) {
	t.Parallel()

	additionOp := shared.Operation{Name: shared.AdditionOperation, SemverType: shared.SemverNoUpdate}
	removalOp := shared.Operation{Name: shared.RemovalOperation, SemverType: shared.SemverNoUpdate}
	sameOp := shared.Operation{Name: shared.NoChangeOperation, SemverType: shared.SemverNoUpdate}
	upgradeMajorOp := shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMajorUpdate}
	upgradeMinorOp := shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverMinorUpdate}
	upgradePatchOp := shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverPatchUpdate}
	downgradeMajorOp := shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMajorUpdate}
	downgradeMinorOp := shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverMinorUpdate}
	downgradePatchOp := shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverPatchUpdate}
	unknownUpdateOp := shared.Operation{Name: shared.UnknownUpdateOperation, SemverType: shared.SemverUnknownUpdate}
	semverExtraUpdateOp := shared.Operation{Name: shared.UnknownUpdateOperation, SemverType: shared.SemverExtraUpdate}

	unmanagedOp := shared.Operation{Name: _testUnmanagedOperation, SemverType: shared.SemverNoUpdate}
	// Following operation (downgrade + semver no update) is not expected to exist
	unmanagedDowngradeOp := shared.Operation{Name: shared.DowngradeOperation, SemverType: shared.SemverNoUpdate}
	// Following operation (upgrade + semver no update) is not expected to exist
	unmanagedUpgradeOp := shared.Operation{Name: shared.UpgradeOperation, SemverType: shared.SemverNoUpdate}

	tests := []struct {
		name      string
		operation shared.Operation
		expected  string
	}{
		{
			name:      "Addition",
			operation: additionOp,
			expected:  "➕️",
		},
		{
			name:      "Removal",
			operation: removalOp,
			expected:  "❌",
		},
		{
			name:      "Same",
			operation: sameOp,
			expected:  "🟰",
		},
		{
			name:      "Major Upgrade",
			operation: upgradeMajorOp,
			expected:  "<sub><sup>🔺.🔹.🔹</sup></sub>",
		},
		{
			name:      "Minor Upgrade",
			operation: upgradeMinorOp,
			expected:  "<sub><sup>🔹.🔺.🔹</sup></sub>",
		},
		{
			name:      "Patch Upgrade",
			operation: upgradePatchOp,
			expected:  "<sub><sup>🔹.🔹.🔺</sup></sub>",
		},
		{
			name:      "Major Downgrade",
			operation: downgradeMajorOp,
			expected:  "<sub><sup>🔻.🔹.🔹</sup></sub>",
		},
		{
			name:      "Minor Downgrade",
			operation: downgradeMinorOp,
			expected:  "<sub><sup>🔹.🔻.🔹</sup></sub>",
		},
		{
			name:      "Patch Downgrade",
			operation: downgradePatchOp,
			expected:  "<sub><sup>🔹.🔹.🔻</sup></sub>",
		},
		{
			name:      "UnknownUpdate",
			operation: unknownUpdateOp,
			expected:  "❓",
		},
		{
			name:      "SemverExtra Update",
			operation: semverExtraUpdateOp,
			expected:  "<sub><sup>🔹.🔹.🔹❓</sup></sub>",
		},
		{
			name:      "Unknown operation",
			operation: unmanagedOp,
			expected:  "❔",
		},
		{
			name:      "Unmanaged upgrade",
			operation: unmanagedUpgradeOp,
			expected:  "❔",
		},
		{
			name:      "Unmanaged downgrade",
			operation: unmanagedDowngradeOp,
			expected:  "❔",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			current := getOperationSymbol(testCase.operation)

			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}
