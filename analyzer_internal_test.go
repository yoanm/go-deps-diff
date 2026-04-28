package depsdiff

import (
	"testing"
)

func TestGuessUpdateOperation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		versionPrevious string
		versionCurrent  string
		wantName        OperationName
		wantSemverType  OperationSemverType
	}{
		{
			name:            "major version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "2.0.0",
			wantName:        UpgradeOperation,
			wantSemverType:  DiffSemverMajor,
		},
		{
			name:            "major version down",
			versionPrevious: "2.0.0",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  DiffSemverMajor,
		},
		{
			name:            "minor version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.1.0",
			wantName:        UpgradeOperation,
			wantSemverType:  DiffSemverMinor,
		},
		{
			name:            "minor version down",
			versionPrevious: "1.1.0",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  DiffSemverMinor,
		},
		{
			name:            "patch version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.1",
			wantName:        UpgradeOperation,
			wantSemverType:  DiffSemverPatch,
		},
		{
			name:            "patch version down",
			versionPrevious: "1.0.1",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  DiffSemverPatch,
		},
		{
			name:            "extra updated",
			versionPrevious: "1.0.0+build.123",
			versionCurrent:  "1.0.0+build.456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  DiffSemverExtra,
		},
		{
			name:            "extra added",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.0+build.123",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  DiffSemverExtra,
		},
		{
			name:            "unparseable version",
			versionPrevious: "abc123",
			versionCurrent:  "def456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  DiffSemverUnknown,
		},
		{
			name:            "one unparseable",
			versionPrevious: "1.0.0",
			versionCurrent:  "def456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  DiffSemverUnknown,
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
