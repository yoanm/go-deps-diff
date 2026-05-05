package summary_test

import (
	"os"
	"testing"

	"github.com/andreyvit/diff"

	"github.com/yoanm/go-deps-diff/shared"
	"github.com/yoanm/go-deps-diff/shared_test"
	"github.com/yoanm/go-deps-diff/summary"
)

func TestIntegration_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		changes    shared.DiffMap
		goldenFile string
	}{
		{
			name:       "Special case - Table with only three column needed",
			changes:    integrationOnlyThreeColumnsNeeded,
			goldenFile: "./testdata/golden-3columns_summary.md",
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
var integrationOnlyThreeColumnsNeeded = shared.DiffMap{
	"caution-dev_only_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-dev_only_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement+dev_req/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"caution-prod_usage-requirement/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "caution-prod_usage-requirement/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"warning-dev_only_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-dev_only_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"warning-prod_usage-transitive/ADDITION+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "warning-prod_usage-transitive/ADDITION+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"important-dev_only_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"important-dev_only_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-dev_only_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"important-prod_usage-requirement+dev_req/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement+dev_req/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"important-prod_usage-requirement/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"important-prod_usage-requirement/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "important-prod_usage-requirement/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"tip-dev_only_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"tip-dev_only_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"tip-dev_only_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-dev_only_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"tip-prod_usage-requirement+dev_req/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement+dev_req/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"tip-prod_usage-requirement/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-requirement/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"tip-prod_usage-transitive/REMOVAL": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.0.1", Label: "3.0.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"tip-prod_usage-transitive/REMOVAL+ABANDONED": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "tip-prod_usage-transitive/REMOVAL+ABANDONED",
			Abandoned:          true,
			Version:            shared.PkgVersion{Raw: "3.2.1", Label: "3.2.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "REMOVAL", SemverType: "NONE"},
	},
	"note-dev_only_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: true,
		},
		Operation: shared.Operation{Name: "NONE", SemverType: "NONE"},
	},
	"note-dev_only_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-dev_only_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            true,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
	"note-prod_usage-requirement/SAME": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-requirement/SAME",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "3.1.1", Label: "3.1.1"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    true,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "NONE", SemverType: "NONE"},
	},
	"note-prod_usage-transitive/ADDITION": {
		Package: &shared_test.TestPkgWrapper{
			Name:               "note-prod_usage-transitive/ADDITION",
			Abandoned:          false,
			Version:            shared.PkgVersion{Raw: "1.18.4", Label: "1.18.4"},
			Link:               "http://www.squizlabs.com/php-codesniffer",
			DevOnly:            false,
			RootRequirement:    false,
			RootDevRequirement: false,
		},
		Operation: shared.Operation{Name: "ADDITION", SemverType: "NONE"},
	},
}
