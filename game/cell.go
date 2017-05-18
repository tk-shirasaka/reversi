package game

import (
	"github.com/tk-shirasaka/ansi"
)

type color int

const (
	BLANK   color = 0x0001
	BLACK   color = 0x0010
	WHITE   color = 0x0100
	PUTABLE color = 0x1000
)

type cell struct {
	color color
}

func (c *cell) string() string {
	switch c.color {
	case BLACK:
		return ansi.Color("● ", ansi.BLACK, ansi.GREEN)
	case WHITE:
		return ansi.Color("● ", ansi.WHITE, ansi.GREEN)
	case PUTABLE:
		return ansi.Color("[]", ansi.RED, ansi.GREEN)
	default:
		return ansi.Color("  ", ansi.WHITE, ansi.GREEN)
	}
}

func (c *cell) change(color color) *cell {
	c.color = color
	return c
}

func (c *cell) is(mask interface{}) bool {
	switch val := mask.(type) {
	case int:
		return (int(c.color) & val) > 0
	case color:
		return (int(c.color) & int(val)) > 0
	default:
		return false
	}
}

func (c *cell) isnot(mask interface{}) bool {
	return c.is(mask) == false
}
