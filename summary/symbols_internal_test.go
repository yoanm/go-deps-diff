package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
)

func Test_getOperationSymbol(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		operation *shared.Operation
		expected  string
	}{
		{name: "Addition", operation: &shared_test.AdditionOp, expected: "➕️"},
		{name: "Removal", operation: &shared_test.RemovalOp, expected: "❌"},
		{name: "Same", operation: &shared_test.SameOp, expected: "🟰"},
		{name: "Major Upgrade", operation: &shared_test.UpgradeMajorOp, expected: "<sub><sup>🔺.🔹.🔹</sup></sub>"},
		{name: "Minor Upgrade", operation: &shared_test.UpgradeMinorOp, expected: "<sub><sup>🔹.🔺.🔹</sup></sub>"},
		{name: "Patch Upgrade", operation: &shared_test.UpgradePatchOp, expected: "<sub><sup>🔹.🔹.🔺</sup></sub>"},
		{name: "Major Downgrade", operation: &shared_test.DowngradeMajorOp, expected: "<sub><sup>🔻.🔹.🔹</sup></sub>"},
		{name: "Minor Downgrade", operation: &shared_test.DowngradeMinorOp, expected: "<sub><sup>🔹.🔻.🔹</sup></sub>"},
		{name: "Patch Downgrade", operation: &shared_test.DowngradePatchOp, expected: "<sub><sup>🔹.🔹.🔻</sup></sub>"},
		{name: "UnknownUpdate", operation: &shared_test.UnknownUpdateOp, expected: "❓"},
		{name: "SemverExtra Update", operation: &shared_test.SemverExtraUpdateOp, expected: "<sub><sup>🔹.🔹.🔹❓</sup></sub>"},
		{name: "Unknown operation", operation: &shared_test.InvalidOp, expected: "❔"},
		{name: "Unmanaged upgrade", operation: &shared_test.InvalidUpgradeOp, expected: "❔"},
		{name: "Unmanaged downgrade", operation: &shared_test.InvalidDowngradeOp, expected: "❔"},
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
