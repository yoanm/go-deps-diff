package markdown

import (
	"iter"
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

func (c *Builder) Details(
	summary string,
	contentBuilder func(c *Builder, indentDepth int),
	opened bool,
	indentDepth int,
) {
	openDirective := ""
	if opened {
		openDirective = " open"
	}

	c.WriteLine("<details"+openDirective+">", indentDepth)
	c.WriteLine("<summary>"+summary+"</summary>", indentDepth+1)
	contentBuilder(c, indentDepth+1)
	c.WriteLine("</details>", indentDepth)
}

func (c *Builder) Header(header string, level int, indentDepth int) {
	c.WriteEol()
	c.WriteLine(strings.Repeat("#", level)+" "+header, indentDepth)
	c.WriteEol()
}

func (c *Builder) HTMLTable(rowIterator iter.Seq[[]string], indentDepth int) {
	c.WriteLine("<table>", indentDepth)

	for cells := range rowIterator {
		c.WriteLine("<tr>"+strings.Join(cells, "")+"</tr>", indentDepth+1)
	}

	c.WriteLine("</table>", indentDepth)
	c.WriteEol()
}
