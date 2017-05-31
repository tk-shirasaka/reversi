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

func (color color) String() string {
	switch color {
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
