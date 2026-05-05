package markdown_test

import (
	"testing"

	"github.com/yoanm/go-deps-diff/summary/markdown"
)

func TestIntegration_Header(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		level       int
		indentDepth int
		expected    string
	}{
		{name: "Level 1", level: 1, indentDepth: 0, expected: "\n# Level 1\n\n"},
		{name: "Level 2", level: 2, indentDepth: 0, expected: "\n## Level 2\n\n"},
		{name: "Level 3", level: 3, indentDepth: 0, expected: "\n### Level 3\n\n"},
		{name: "Level 4", level: 4, indentDepth: 0, expected: "\n#### Level 4\n\n"},
		{name: "Level 5", level: 5, indentDepth: 0, expected: "\n##### Level 5\n\n"},
		{name: "Level 1 - indented", level: 1, indentDepth: 3, expected: "\n      # Level 1 - indented\n\n"},
		{name: "Level 2 - indented", level: 2, indentDepth: 3, expected: "\n      ## Level 2 - indented\n\n"},
		{name: "Level 3 - indented", level: 3, indentDepth: 3, expected: "\n      ### Level 3 - indented\n\n"},
		{name: "Level 4 - indented", level: 4, indentDepth: 3, expected: "\n      #### Level 4 - indented\n\n"},
		{name: "Level 5 - indented", level: 5, indentDepth: 3, expected: "\n      ##### Level 5 - indented\n\n"},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			builder := markdown.NewBuilder()
			builder.Header(testCase.name, testCase.level, testCase.indentDepth)

			current := builder.String()
			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}

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

func TestIntegration_Details(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		summary     string
		content     string
		opened      bool
		indentDepth int
		expected    string
	}{
		{
			name:        "Base",
			summary:     "SUMMARY",
			content:     "__CONTENT__\n",
			opened:      false,
			indentDepth: 0,
			expected:    "<details>\n  <summary>SUMMARY</summary>\n__CONTENT__\n</details>\n",
		},
		{
			name:        "Opened",
			summary:     "SUMMARY",
			content:     "__CONTENT__\n",
			opened:      true,
			indentDepth: 0,
			expected:    "<details open>\n  <summary>SUMMARY</summary>\n__CONTENT__\n</details>\n",
		},
		{
			name:        "Indented",
			summary:     "SUMMARY",
			content:     "__CONTENT__\n",
			opened:      false,
			indentDepth: 3,
			expected:    "      <details>\n        <summary>SUMMARY</summary>\n__CONTENT__\n      </details>\n",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			builder := markdown.NewBuilder()
			builder.Details(
				testCase.summary,
				func(c *markdown.Builder, indentDepth int) { c.Write(testCase.content) },
				testCase.opened,
				testCase.indentDepth,
			)

			current := builder.String()
			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}
