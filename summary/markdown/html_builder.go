package markdown

import (
	"iter"
	"strings"
)

func (c *Builder) HTMLTable(rowIterator iter.Seq[[]string], indentDepth int) {
	c.WriteLine("<table>", indentDepth)

	for cells := range rowIterator {
		c.WriteLine("<tr>"+strings.Join(cells, "")+"</tr>", indentDepth+1)
	}

	c.WriteLine("</table>", indentDepth)
	c.WriteEol()
}
