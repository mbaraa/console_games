package tui

type TermColor string

const (
	ColorReset     TermColor = "\033[0m"
	ColorBoldGreen TermColor = "\033[1m\033[32m"
	ColorBoldRed   TermColor = "\033[1m\033[31m"
)
