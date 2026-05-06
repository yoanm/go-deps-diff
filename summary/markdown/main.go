package markdown

import (
	"strings"
)

func NewBuilder() *Builder {
	return &Builder{
		builder:      &strings.Builder{},
		eolString:    "\n",
		indentString: "  ",
	}
}

type Builder struct {
	builder      *strings.Builder
	eolString    string
	indentString string
}
