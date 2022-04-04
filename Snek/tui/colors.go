package tui

type TermColor string

const (
	ColorReset     TermColor = "\033[0m"
	ColorBoldGreen TermColor = "\033[1m\033[32m"
	ColorBoldRed   TermColor = "\033[1m\033[31m"
)

func (c TermColor) StringColored(s string) string {
	return string(c) + s + string(ColorReset)
}

type Cell interface {
	Position() Point2
	Name() string
	Color() TermColor
	Content() string
	String() string
}

type ColoredCell struct {
	name     string
	content  string
	color    TermColor
	position Point2
}

func NewColoredCell(name, content string, color TermColor, position Point2) *ColoredCell {
	return &ColoredCell{
		name:     name,
		content:  content,
		color:    color,
		position: position,
	}
}

func (c *ColoredCell) Position() Point2 {
	return c.position
}

func (c *ColoredCell) Name() string {
	return c.name
}

func (c *ColoredCell) Content() string {
	return c.content
}

func (c *ColoredCell) Color() TermColor {
	return c.color
}

func (c ColoredCell) String() string {
	return c.color.StringColored(c.content)
}
