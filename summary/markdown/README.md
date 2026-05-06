# markdown

## Types

### type [Builder](./main.go#L15)

`type Builder struct { ... }`

#### func [NewBuilder](./main.go#L7)

`func NewBuilder() *Builder`

#### func (*Builder) [Details](./markdown_builder.go#L5)

`func (c *Builder) Details(
    summary string,
    contentBuilder func(c *Builder, indentDepth int),
    opened bool,
    indentDepth int,
)`

#### func (*Builder) [HTMLTable](./html_builder.go#L8)

`func (c *Builder) HTMLTable(rowIterator iter.Seq[[]string], indentDepth int)`

#### func (*Builder) [Header](./markdown_builder.go#L22)

`func (c *Builder) Header(header string, level int, indentDepth int)`

#### func (*Builder) [String](./builder.go#L7)

`func (c *Builder) String() string`

#### func (*Builder) [Write](./builder.go#L21)

`func (c *Builder) Write(v string)`

#### func (*Builder) [WriteEol](./builder.go#L17)

`func (c *Builder) WriteEol()`

#### func (*Builder) [WriteLine](./builder.go#L11)

`func (c *Builder) WriteLine(line string, indentDepth int)`

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
