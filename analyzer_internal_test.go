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
		wantDirection   OperationDirection
	}{
		{
			name:            "major version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "2.0.0",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverMajor,
			wantDirection:   DiffDirectionUp,
		},
		{
			name:            "major version down",
			versionPrevious: "2.0.0",
			versionCurrent:  "1.0.0",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverMajor,
			wantDirection:   DiffDirectionDown,
		},
		{
			name:            "minor version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.1.0",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverMinor,
			wantDirection:   DiffDirectionUp,
		},
		{
			name:            "minor version down",
			versionPrevious: "1.1.0",
			versionCurrent:  "1.0.0",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverMinor,
			wantDirection:   DiffDirectionDown,
		},
		{
			name:            "patch version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.1",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverPatch,
			wantDirection:   DiffDirectionUp,
		},
		{
			name:            "patch version down",
			versionPrevious: "1.0.1",
			versionCurrent:  "1.0.0",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverPatch,
			wantDirection:   DiffDirectionDown,
		},
		{
			name:            "extra updated",
			versionPrevious: "1.0.0+build.123",
			versionCurrent:  "1.0.0+build.456",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverExtra,
			wantDirection:   DiffDirectionUnknown,
		},
		{
			name:            "extra added",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.0+build.123",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverExtra,
			wantDirection:   DiffDirectionUnknown,
		},
		{
			name:            "unparseable version",
			versionPrevious: "abc123",
			versionCurrent:  "def456",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverUnknown,
			wantDirection:   DiffDirectionUnknown,
		},
		{
			name:            "one unparseable",
			versionPrevious: "1.0.0",
			versionCurrent:  "def456",
			wantName:        UpdatedPackage,
			wantSemverType:  DiffSemverUnknown,
			wantDirection:   DiffDirectionUnknown,
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

			if result.Direction != testData.wantDirection {
				t.Errorf("guessUpdateOperation() Direction = %s, want %s", result.Direction, testData.wantDirection)
			}
		})
	}
}
