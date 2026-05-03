package depsdiff

import (
	"testing"

	"github.com/yoanm/go-deps-diff/shared"
)

func TestGuessUpdateOperation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		versionPrevious string
		versionCurrent  string
		wantName        shared.OperationName
		wantSemverType  shared.OperationSemverType
	}{
		{
			name:            "major version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "2.0.0",
			wantName:        shared.UpgradeOperation,
			wantSemverType:  shared.SemverMajorUpdate,
		},
		{
			name:            "major version down",
			versionPrevious: "2.0.0",
			versionCurrent:  "1.0.0",
			wantName:        shared.DowngradeOperation,
			wantSemverType:  shared.SemverMajorUpdate,
		},
		{
			name:            "minor version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.1.0",
			wantName:        shared.UpgradeOperation,
			wantSemverType:  shared.SemverMinorUpdate,
		},
		{
			name:            "minor version down",
			versionPrevious: "1.1.0",
			versionCurrent:  "1.0.0",
			wantName:        shared.DowngradeOperation,
			wantSemverType:  shared.SemverMinorUpdate,
		},
		{
			name:            "patch version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.1",
			wantName:        shared.UpgradeOperation,
			wantSemverType:  shared.SemverPatchUpdate,
		},
		{
			name:            "patch version down",
			versionPrevious: "1.0.1",
			versionCurrent:  "1.0.0",
			wantName:        shared.DowngradeOperation,
			wantSemverType:  shared.SemverPatchUpdate,
		},
		{
			name:            "extra updated",
			versionPrevious: "1.0.0+build.123",
			versionCurrent:  "1.0.0+build.456",
			wantName:        shared.UnknownUpdateOperation,
			wantSemverType:  shared.SemverExtraUpdate,
		},
		{
			name:            "extra added",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.0+build.123",
			wantName:        shared.UnknownUpdateOperation,
			wantSemverType:  shared.SemverExtraUpdate,
		},
		{
			name:            "unparseable version",
			versionPrevious: "abc123",
			versionCurrent:  "def456",
			wantName:        shared.UnknownUpdateOperation,
			wantSemverType:  shared.SemverUnknownUpdate,
		},
		{
			name:            "one unparseable",
			versionPrevious: "1.0.0",
			versionCurrent:  "def456",
			wantName:        shared.UnknownUpdateOperation,
			wantSemverType:  shared.SemverUnknownUpdate,
		},
	}

	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			result := guessUpdateOperation(testData.versionPrevious, testData.versionCurrent)
			if result.Name != testData.wantName {
				t.Errorf("guessUpdateOperation() Name = %s, want %s", result.Name, testData.wantName)
			}

			if result.SemverType != testData.wantSemverType {
				t.Errorf("guessUpdateOperation() SemverType = %s, want %s", result.SemverType, testData.wantSemverType)
			}
		})
	}
}
