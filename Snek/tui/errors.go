package tui

import "errors"

var (
	ErrPlaneOverflow  = errors.New("coordinates overflows the given plane")
	ErrPlaneUnderflow = errors.New("coordinates underflows the given plane")
)
