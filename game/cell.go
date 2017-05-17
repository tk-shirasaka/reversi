package game

import (
	"github.com/tk-shirasaka/ansi"
)

type color int

const (
	BLANK color = iota
	BLACK
	WHITE
	PUTABLE
)

type cell struct {
	color color
}

func (c *cell) string() string {
	switch c.color {
	case BLACK:
		return ansi.Ansi("● ", ansi.BLACK, ansi.GREEN)
	case WHITE:
		return ansi.Ansi("● ", ansi.WHITE, ansi.GREEN)
	case PUTABLE:
		return ansi.Ansi("[]", ansi.RED, ansi.GREEN)
	default:
		return ansi.Ansi("  ", ansi.WHITE, ansi.GREEN)
	}
}

func (c *cell) blank() *cell {
	c.color = BLANK
	return c
}

func (c *cell) black() *cell {
	c.color = BLACK
	return c
}

func (c *cell) white() *cell {
	c.color = WHITE
	return c
}

func (c *cell) putable() *cell {
	c.color = PUTABLE
	return c
}
