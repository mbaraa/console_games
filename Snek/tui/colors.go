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

type ColoredCell struct {
	content string
	color   TermColor
}

func NewCell(content string, color TermColor) ColoredCell {
	return ColoredCell{
		content: content,
		color:   color,
	}
}

func (c ColoredCell) Content() string {
	return c.content
}

func (c ColoredCell) Color() TermColor {
	return c.color
}

func (c ColoredCell) String() string {
	return c.color.StringColored(c.content)
}
