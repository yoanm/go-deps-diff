package summary_test

import (
	"os"
	"testing"

	"github.com/andreyvit/diff"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
	"github.com/yoanm/go-deps-diff/summary"
)

func TestIntegration_GenerateForChanges(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		changes    shared.DiffMap
		goldenFile string
	}{
		{
			name:       "Full",
			changes:    _integrationFullChanges,
			goldenFile: "./testdata/golden-full-summary.md",
		},
		{
			name:       "Special case - shortest table size - Table with only three column needed",
			changes:    _integrationOnlyThreeColumnsNeeded,
			goldenFile: "./testdata/golden-3columns-summary.md",
		},
		{
			name:       "Special case - force opened/closed details",
			changes:    _integrationForceOpenedClosed,
			goldenFile: "./testdata/golden-force_opened_details.md",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			expected, err := os.ReadFile(testCase.goldenFile)
			if err != nil {
				t.Errorf("Diff() error while reading previous lock file = %v", err)

				return
			}

			current := summary.GenerateForChanges(testCase.changes)

			if string(expected) != current {
				t.Errorf("unexpected output: diff %s", diff.LineDiff(string(expected), current))
			}
		})
	}
}

//nolint:gochecknoglobals // Just to keep it outside the function
var _integrationFullChanges = shared.DiffMap{
	"caution-dev_only_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"caution-dev_only_usage-requirement/SEMVER_MAJOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/SEMVER_MAJOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.2.0", Label: "1.2.0", Semver: &shared.SemverVersion{Major: 1, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2", Label: "2.9.2", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"caution-dev_only_usage-requirement/UNKNOWN_UPDATE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/UNKNOWN_UPDATE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "12345678", Label: "2.9.x-dev#1234567", Semver: nil}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UnknownUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2acf168", Label: "2.9.x-dev#2acf168", Semver: nil}, //nolint:lll // Meaningless for tests !
	},
	"caution-dev_only_usage-requirement/UNKNOWN_UPDATE+SEMVER_EXTRA": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/UNKNOWN_UPDATE+SEMVER_EXTRA",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "2.9.2+alpha", Label: "2.9.2+alpha", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.SemverExtraUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2+beta", Label: "2.9.2+beta", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: "+beta"}}, //nolint:lll // Meaningless for tests !
	},
	"caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"caution-prod_usage-requirement+dev_req/SEMVER_MAJOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/SEMVER_MAJOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2", Label: "2.9.2", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"caution-prod_usage-requirement+dev_req/UNKNOWN_UPDATE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/UNKNOWN_UPDATE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "12345678", Label: "2.9.x-dev#1234567", Semver: nil}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UnknownUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2acf168", Label: "2.9.x-dev#2acf168", Semver: nil}, //nolint:lll // Meaningless for tests !
	},
	"caution-prod_usage-requirement+dev_req/UNKNOWN_UPDATE+SEMVER_EXTRA": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/UNKNOWN_UPDATE+SEMVER_EXTRA",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "2.9.2+alpha", Label: "2.9.2+alpha", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.SemverExtraUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2+beta", Label: "2.9.2+beta", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: "+beta"}}, //nolint:lll // Meaningless for tests !
	},
	"caution-prod_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"caution-prod_usage-requirement/SEMVER_MAJOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/SEMVER_MAJOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2", Label: "2.9.2", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"caution-prod_usage-requirement/UNKNOWN_UPDATE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/UNKNOWN_UPDATE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "12345678", Label: "2.9.x-dev#1234567", Semver: nil}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UnknownUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2acf168", Label: "2.9.x-dev#2acf168", Semver: nil}, //nolint:lll // Meaningless for tests !
	},
	"caution-prod_usage-requirement/UNKNOWN_UPDATE+SEMVER_EXTRA": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/UNKNOWN_UPDATE+SEMVER_EXTRA",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "2.9.2+alpha", Label: "2.9.2+alpha", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.SemverExtraUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2+beta", Label: "2.9.2+beta", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: "+beta"}}, //nolint:lll // Meaningless for tests !
	},
	"warning-dev_only_usage-requirement/SEMVER_MAJOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-requirement/SEMVER_MAJOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-dev_only_usage-requirement/ADDITION_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-requirement/ADDITION_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-dev_only_usage-requirement/SEMVER_MINOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-requirement/SEMVER_MINOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-dev_only_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-dev_only_usage-transitive/ADDITION_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/ADDITION_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-dev_only_usage-transitive/SEMVER_MAJOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/SEMVER_MAJOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.2.0", Label: "1.2.0", Semver: &shared.SemverVersion{Major: 1, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2", Label: "2.9.2", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-dev_only_usage-transitive/UNKNOWN_UPDATE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/UNKNOWN_UPDATE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "12345678", Label: "2.9.x-dev#1234567", Semver: nil}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UnknownUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2acf168", Label: "2.9.x-dev#2acf168", Semver: nil}, //nolint:lll // Meaningless for tests !,
	},
	"warning-dev_only_usage-transitive/UNKNOWN_UPDATE+SEMVER_EXTRA": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/UNKNOWN_UPDATE+SEMVER_EXTRA",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "2.9.2+alpha", Label: "2.9.2+alpha", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.SemverExtraUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2+beta", Label: "2.9.2+beta", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: "+beta"}}, //nolint:lll // Meaningless for tests !
	},
	"warning-prod_usage-requirement+dev_req/SEMVER_MAJOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement+dev_req/SEMVER_MAJOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-requirement+dev_req/SEMVER_MINOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement+dev_req/SEMVER_MINOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-requirement/SEMVER_MAJOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement/SEMVER_MAJOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-requirement/SEMVER_MINOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement/SEMVER_MINOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-requirement/ADDITION_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement/ADDITION_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.AdditionOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-requirement+dev_req/ADDITION_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-requirement+dev_req/ADDITION_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.AdditionOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-prod_usage-transitive/ADDITION_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/ADDITION_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-prod_usage-transitive/SEMVER_MAJOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/SEMVER_MAJOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.2.0", Label: "1.2.0", Semver: &shared.SemverVersion{Major: 1, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2", Label: "2.9.2", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"warning-prod_usage-transitive/UNKNOWN_UPDATE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/UNKNOWN_UPDATE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "12345678", Label: "2.9.x-dev#1234567", Semver: nil}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UnknownUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2acf168", Label: "2.9.x-dev#2acf168", Semver: nil}, //nolint:lll // Meaningless for tests !
	},
	"warning-prod_usage-transitive/UNKNOWN_UPDATE+SEMVER_EXTRA": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/UNKNOWN_UPDATE+SEMVER_EXTRA",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "2.9.2+alpha", Label: "2.9.2+alpha", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: ""}}, //nolint:lll // Meaningless for tests !
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.SemverExtraUpdateOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.2+beta", Label: "2.9.2+beta", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 2, Extra: "+beta"}}, //nolint:lll // Meaningless for tests !
	},
	"important-dev_only_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-dev_only_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-dev_only_usage-requirement/SEMVER_PATCH_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/SEMVER_PATCH_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-dev_only_usage-requirement/SAME_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/SAME_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"important-dev_only_usage-transitive/SEMVER_MAJOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-transitive/SEMVER_MAJOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-dev_only_usage-transitive/SEMVER_MINOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-transitive/SEMVER_MINOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-prod_usage-requirement+dev_req/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement+dev_req/SEMVER_PATCH_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/SEMVER_PATCH_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.DowngradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-prod_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement/SEMVER_PATCH_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/SEMVER_PATCH_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-prod_usage-requirement/SAME_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/SAME_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"important-prod_usage-requirement+dev_req/SAME_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/SAME_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"important-prod_usage-transitive/SEMVER_MAJOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-transitive/SEMVER_MAJOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMajorOp,
		PreviousVersion: shared.PkgVersion{Raw: "2.9.3", Label: "2.9.3", Semver: &shared.SemverVersion{Major: 2, Minor: 9, Patch: 3, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"important-prod_usage-transitive/SEMVER_MINOR_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-transitive/SEMVER_MINOR_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-dev_only_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-dev_only_usage-requirement/SEMVER_MINOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-requirement/SEMVER_MINOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-dev_only_usage-requirement/SEMVER_MINOR_UPGRADE+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-requirement/SEMVER_MINOR_UPGRADE+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.0", Label: "3.2.0", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-dev_only_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-dev_only_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-dev_only_usage-transitive/SEMVER_PATCH_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/SEMVER_PATCH_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-dev_only_usage-transitive/SAME_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/SAME_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"tip-prod_usage-requirement+dev_req/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement+dev_req/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-prod_usage-requirement+dev_req/SEMVER_MINOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement+dev_req/SEMVER_MINOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-prod_usage-requirement+dev_req/SEMVER_MINOR_UPGRADE+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement+dev_req/SEMVER_MINOR_UPGRADE+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.0", Label: "3.2.0", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-prod_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-prod_usage-requirement/SEMVER_MINOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement/SEMVER_MINOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-prod_usage-requirement/SEMVER_MINOR_UPGRADE+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement/SEMVER_MINOR_UPGRADE+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.0", Label: "3.2.0", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-prod_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-prod_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-prod_usage-transitive/SEMVER_PATCH_DOWNGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/SEMVER_PATCH_DOWNGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.DowngradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"tip-prod_usage-transitive/SAME_NOT_SEMVER": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/SAME_NOT_SEMVER",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "abcdefghijk", Label: "dev-master#abcdefgh", Semver: nil},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.SameOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-dev_only_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-requirement/SAME+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-requirement/SEMVER_PATCH_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SEMVER_PATCH_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-dev_only_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"note-dev_only_usage-transitive/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-transitive/SAME+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/SAME+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-transitive/SEMVER_MINOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/SEMVER_MINOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-dev_only_usage-transitive/SEMVER_MINOR_UPGRADE+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/SEMVER_MINOR_UPGRADE+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.0", Label: "3.2.0", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-dev_only_usage-transitive/SEMVER_PATCH_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/SEMVER_PATCH_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-prod_usage-requirement+dev_req/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement+dev_req/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-requirement+dev_req/SAME+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement+dev_req/SAME+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-requirement+dev_req/SEMVER_PATCH_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement+dev_req/SEMVER_PATCH_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation:       shared_test.UpgradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-prod_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-requirement/SAME+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SAME+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-requirement/SEMVER_PATCH_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SEMVER_PATCH_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-prod_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"note-prod_usage-transitive/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-transitive/SAME+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/SAME+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-transitive/SEMVER_MINOR_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/SEMVER_MINOR_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.0", Label: "3.0.0", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-prod_usage-transitive/SEMVER_MINOR_UPGRADE+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/SEMVER_MINOR_UPGRADE+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.0", Label: "3.2.0", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradeMinorOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
	"note-prod_usage-transitive/SEMVER_PATCH_UPGRADE": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/SEMVER_PATCH_UPGRADE",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation:       shared_test.UpgradePatchOp,
		PreviousVersion: shared.PkgVersion{Raw: "3.1.0", Label: "3.1.0", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 0, Extra: ""}}, //nolint:lll // Meaningless for tests !,
	},
}

//nolint:gochecknoglobals // Just to keep it outside the function
var _integrationOnlyThreeColumnsNeeded = shared.DiffMap{
	"caution-dev_only_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"caution-prod_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-dev_only_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"warning-prod_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"important-dev_only_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-dev_only_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement+dev_req/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"important-prod_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-dev_only_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-dev_only_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-dev_only_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-prod_usage-requirement+dev_req/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement+dev_req/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-prod_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"tip-prod_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1", Semver: &shared.SemverVersion{Major: 3, Minor: 0, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"tip-prod_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1", Semver: &shared.SemverVersion{Major: 3, Minor: 2, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.RemovalOp,
	},
	"note-dev_only_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
	"note-prod_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4", Semver: &shared.SemverVersion{Major: 1, Minor: 18, Patch: 4, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared_test.AdditionOp,
	},
}

//nolint:gochecknoglobals // Just to keep it outside the function
var _integrationForceOpenedClosed = shared.DiffMap{
	"note-dev_only_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-requirement/SAME-2": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME-2",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-requirement/SAME-3": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME-3",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-dev_only_usage-requirement/SAME-4": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME-4",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared_test.SameOp,
	},
	"note-prod_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1", Semver: &shared.SemverVersion{Major: 3, Minor: 1, Patch: 1, Extra: ""}}, //nolint:lll // Meaningless for tests !,
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared_test.SameOp,
	},
}
