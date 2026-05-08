package depsdiff_test

import (
	"os"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/contract"
)

func TestIntegration_Fixtures(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		reqFilePath  string
		lockFilePath string
	}{
		{
			name:         "Simple",
			reqFilePath:  "./managers/composer/testdata/composer-simple.json",
			lockFilePath: "./managers/composer/testdata/composer-simple.lock",
		},
		{
			name:         "Complex",
			reqFilePath:  "./managers/composer/testdata/composer-complex.json",
			lockFilePath: "./managers/composer/testdata/composer-complex.lock",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			reqContent, err := os.ReadFile(testCase.reqFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading requirement file = %v", err)

				return
			}

			lockContent, err := os.ReadFile(testCase.lockFilePath)
			if err != nil {
				t.Errorf("Diff() error while reading lock file = %v", err)

				return
			}

			out, err := depsdiff.ComposerDiff(
				&depsdiff.PkgManagerInput{
					Lock:        lockContent,
					Requirement: reqContent,
				},
				&depsdiff.PkgManagerInput{
					Lock:        lockContent,
					Requirement: reqContent,
				},
			)
			if err != nil {
				t.Errorf("Diff() error = %v", err)

				return
			}

			if len(out) == 0 {
				t.Fatal("Diff() result check failed: no packages found in the output")
			}

			for _, change := range out {
				if change.Operation.Name != contract.NoChangeOperation {
					t.Errorf(
						"Diff() result check failed: expected all packages to be unchanged, but %s isn't",
						change.Package.GetName(),
					)
				}
			}
		})
	}
}
