package depsdiff

import (
	"testing"

	"github.com/yoanm/go-deps-diff/contract"
	"github.com/yoanm/go-deps-diff/contract/semver"
	difftesting "github.com/yoanm/go-deps-diff/testing"
)

func TestGuessUpdateOperation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		previous contract.PkgVersion
		current  contract.PkgVersion
		expected contract.Operation
	}{
		{
			name:     "major version up",
			previous: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "2.0.0", Label: "2.0.0", Semver: &semver.Version{Major: 2, Minor: 0, Patch: 0, Extra: ""}},
			expected: difftesting.UpgradeMajorOp,
		},
		{
			name:     "major version down",
			previous: contract.PkgVersion{Raw: "2.0.0", Label: "2.0.0", Semver: &semver.Version{Major: 2, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			expected: difftesting.DowngradeMajorOp,
		},
		{
			name:     "minor version up",
			previous: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "1.1.0", Label: "1.1.0", Semver: &semver.Version{Major: 1, Minor: 1, Patch: 0, Extra: ""}},
			expected: difftesting.UpgradeMinorOp,
		},
		{
			name:     "minor version down",
			previous: contract.PkgVersion{Raw: "1.1.0", Label: "1.1.0", Semver: &semver.Version{Major: 1, Minor: 1, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			expected: difftesting.DowngradeMinorOp,
		},
		{
			name:     "patch version up",
			previous: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "1.0.1", Label: "1.0.1", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 1, Extra: ""}},
			expected: difftesting.UpgradePatchOp,
		},
		{
			name:     "patch version down",
			previous: contract.PkgVersion{Raw: "1.0.1", Label: "1.0.1", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 1, Extra: ""}},
			current:  contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			expected: difftesting.DowngradePatchOp,
		},
		{
			name:     "extra updated",
			previous: contract.PkgVersion{Raw: "abcdefghij", Label: "1.0.0+build.123#abcdefgh", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: "+build.123"}},
			current:  contract.PkgVersion{Raw: "abcdefghij", Label: "1.0.0+build.456#abcdefgh", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: "+build.456"}},
			expected: difftesting.SemverExtraUpdateOp,
		},
		{
			name:     "extra added",
			previous: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "abcdefghij", Label: "1.0.0+build.123#abcdefgh", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: "+build.123"}},
			expected: difftesting.SemverExtraUpdateOp,
		},
		{
			name:     "non-semver versions",
			previous: contract.PkgVersion{Raw: "abcdefghij", Label: "dev-master#abcdefh", Semver: nil},
			current:  contract.PkgVersion{Raw: "klmnopqrs", Label: "dev-master#klmnopq", Semver: nil},
			expected: difftesting.UnknownUpdateOp,
		},
		{
			name:     "current as non-semver version",
			previous: contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			current:  contract.PkgVersion{Raw: "klmnopqrs", Label: "dev-master#klmnopq", Semver: nil},
			expected: difftesting.UnknownUpdateOp,
		},
		{
			name:     "previous as non-semver version",
			previous: contract.PkgVersion{Raw: "klmnopqrs", Label: "dev-master#klmnopq", Semver: nil},
			current:  contract.PkgVersion{Raw: "1.0.0", Label: "1.0.0", Semver: &semver.Version{Major: 1, Minor: 0, Patch: 0, Extra: ""}},
			expected: difftesting.UnknownUpdateOp,
		},
	}

	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			result := guessUpdateOperation(testData.previous, testData.current)

			if err := difftesting.ValidateOperation(result, testData.expected); err != nil {
				t.Error(err)
			}
		})
	}
}
