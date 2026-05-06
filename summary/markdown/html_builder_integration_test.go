package markdown_test

import (
	"testing"

	"github.com/yoanm/go-deps-diff/summary/markdown"
)

func TestIntegration_HTMLTable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		cells       []string
		indentDepth int
		expected    string
	}{
		{
			name:        "Base",
			cells:       []string{"<td>Cell1</td>", "<td>Cell2</td>", "<td>Cell3</td>"},
			indentDepth: 0,
			expected:    "<table>\n  <tr><td>Cell1</td><td>Cell2</td><td>Cell3</td></tr>\n</table>\n\n",
		},
		{
			name:        "Indented",
			cells:       []string{"<td>Cell1</td>", "<td>Cell2</td>", "<td>Cell3</td>"},
			indentDepth: 3,
			expected:    "      <table>\n        <tr><td>Cell1</td><td>Cell2</td><td>Cell3</td></tr>\n      </table>\n\n",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			builder := markdown.NewBuilder()
			builder.HTMLTable(
				func(yield func([]string) bool) { yield(testCase.cells) },
				testCase.indentDepth,
			)

			current := builder.String()
			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}
