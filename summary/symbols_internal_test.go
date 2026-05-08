package summary

import (
	"testing"

	"github.com/yoanm/go-deps-diff/contract"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func Test_getOperationSymbol(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		operation contract.Operation
		expected  string
	}{
		{name: "Addition", operation: difftesting.AdditionOp, expected: "➕️"},
		{name: "Removal", operation: difftesting.RemovalOp, expected: "❌"},
		{name: "Same", operation: difftesting.SameOp, expected: "🟰"},
		{name: "Major Upgrade", operation: difftesting.UpgradeMajorOp, expected: "<sub><sup>🔺.🔹.🔹</sup></sub>"},
		{name: "Minor Upgrade", operation: difftesting.UpgradeMinorOp, expected: "<sub><sup>🔹.🔺.🔹</sup></sub>"},
		{name: "Patch Upgrade", operation: difftesting.UpgradePatchOp, expected: "<sub><sup>🔹.🔹.🔺</sup></sub>"},
		{name: "Major Downgrade", operation: difftesting.DowngradeMajorOp, expected: "<sub><sup>🔻.🔹.🔹</sup></sub>"},
		{name: "Minor Downgrade", operation: difftesting.DowngradeMinorOp, expected: "<sub><sup>🔹.🔻.🔹</sup></sub>"},
		{name: "Patch Downgrade", operation: difftesting.DowngradePatchOp, expected: "<sub><sup>🔹.🔹.🔻</sup></sub>"},
		{name: "UnknownUpdate", operation: difftesting.UnknownUpdateOp, expected: "❓"},
		{name: "SemverExtra Update", operation: difftesting.SemverExtraUpdateOp, expected: "<sub><sup>🔹.🔹.🔹❓</sup></sub>"},
		{name: "Unknown operation", operation: difftesting.InvalidOp, expected: "❔"},
		{name: "Unmanaged upgrade", operation: difftesting.InvalidUpgradeOp, expected: "❔"},
		{name: "Unmanaged downgrade", operation: difftesting.InvalidDowngradeOp, expected: "❔"},
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
