package markdown

import "strings"

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
