package markdown_test

import (
	"testing"

	"github.com/yoanm/go-deps-diff/summary/markdown"
)

func TestIntegration_WriteEol(t *testing.T) {
	t.Parallel()

	builder := markdown.NewBuilder()
	builder.WriteEol()

	current := builder.String()
	expected := "\n"

	if expected != current {
		t.Fatalf("unexpected output: got %s, want %s", current, expected)
	}
}

func TestIntegration_WriteLine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		indentDepth int
		expected    string
	}{
		{name: "Base", indentDepth: 0, expected: "Base\n"},
		{name: "Indented", indentDepth: 3, expected: "      Indented\n"},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			builder := markdown.NewBuilder()
			builder.WriteLine(testCase.name, testCase.indentDepth)

			current := builder.String()
			if testCase.expected != current {
				t.Errorf("unexpected output: got %s, want %s", current, testCase.expected)
			}
		})
	}
}
