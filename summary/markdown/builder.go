package markdown

import (
	"strings"
)

func (c *Builder) String() string {
	return c.builder.String()
}

func (c *Builder) WriteLine(line string, indentDepth int) {
	c.Write(strings.Repeat(c.indentString, indentDepth))
	c.Write(line)
	c.WriteEol()
}

func (c *Builder) WriteEol() {
	c.Write(c.eolString)
}

func (c *Builder) Write(v string) {
	_, err := c.builder.WriteString(v)
	if nil != err {
		panic(err)
	}
}
