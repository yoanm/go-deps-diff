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
			wantSemverType:  SemverMajorUpdate,
		},
		{
			name:            "major version down",
			versionPrevious: "2.0.0",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  SemverMajorUpdate,
		},
		{
			name:            "minor version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.1.0",
			wantName:        UpgradeOperation,
			wantSemverType:  SemverMinorUpdate,
		},
		{
			name:            "minor version down",
			versionPrevious: "1.1.0",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  SemverMinorUpdate,
		},
		{
			name:            "patch version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.1",
			wantName:        UpgradeOperation,
			wantSemverType:  SemverPatchUpdate,
		},
		{
			name:            "patch version down",
			versionPrevious: "1.0.1",
			versionCurrent:  "1.0.0",
			wantName:        DowngradeOperation,
			wantSemverType:  SemverPatchUpdate,
		},
		{
			name:            "extra updated",
			versionPrevious: "1.0.0+build.123",
			versionCurrent:  "1.0.0+build.456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  SemverExtraUpdate,
		},
		{
			name:            "extra added",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.0+build.123",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  SemverExtraUpdate,
		},
		{
			name:            "unparseable version",
			versionPrevious: "abc123",
			versionCurrent:  "def456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  SemverUnknownUpdate,
		},
		{
			name:            "one unparseable",
			versionPrevious: "1.0.0",
			versionCurrent:  "def456",
			wantName:        UnknownUpdateOperation,
			wantSemverType:  SemverUnknownUpdate,
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
