package game

import (
	"strconv"
)

type field struct {
	turn  color
	dirty bool
	cells [8][8]cell
}

func Init() *field {
	f := new(field)
	f.turn = BLACK
	f.dirty = true
	f.cells[3][3].black()
	f.cells[3][4].white()
	f.cells[4][3].white()
	f.cells[4][4].black()

	return f
}

func (f *field) itelator(i int, j int, i_offset int, j_offset int, callback func(*cell) bool) {
	for i >= 0 && i <= 7 && j >= 0 && j <= 7 {
		if callback(&f.cells[i][j]) == false {
			break
		}
		i += i_offset
		j += j_offset
	}
}

func (f *field) check() {
	if f.dirty == false {
		return
	}
	for i, line := range f.cells {
		for j, _ := range line {
			c := &f.cells[i][j]
			if c.color != BLACK && c.color != WHITE {
				c.blank()
			}
		}
	}
	for i, line := range f.cells {
		for j, _ := range line {
			step := 0
			checker := func(c *cell) bool {
				if (c.color == BLACK || c.color == WHITE) && step > 0 {
					if step == 1 && c.color != f.turn {
						step++
					} else if step == 2 && c.color == f.turn {
						step = 0
						f.cells[i][j].putable()
						return false
					}
				} else if !(c.color == BLACK || c.color == WHITE) && step == 0 {
					step++
				} else {
					step = 0
					return false
				}
				return true
			}

			f.itelator(i, j, -1, -1, checker)
			f.itelator(i, j, -1, 0, checker)
			f.itelator(i, j, -1, 1, checker)
			f.itelator(i, j, 0, -1, checker)
			f.itelator(i, j, 0, 1, checker)
			f.itelator(i, j, 1, -1, checker)
			f.itelator(i, j, 1, 0, checker)
			f.itelator(i, j, 1, 1, checker)
		}
	}
}

func (f *field) Select(j int, i int) {
	if i < 1 || i > 8 || j < 1 || j > 8 {
		return
	}
	i--
	j--
	f.check()
	if f.cells[i][j].color == PUTABLE {
		step := 0
		next := WHITE
		if f.turn == next {
			next = BLACK
		}
		selector := func(c *cell) bool {
			if step == 0 {
				c.color = f.turn
				step = 1
			} else if step == 1 && c.color == next {
				step = 2
			} else if step == 2 && c.color == f.turn {
				step = 3
				return false
			} else if step == 3 {
				step = 4
			} else if step == 4 && c.color == next {
				step = 5
				c.color = f.turn
			} else {
				step = 0
				return false
			}
			return true
		}
		f.itelator(i, j, -1, -1, selector)
		f.itelator(i, j, -1, -1, selector)
		f.itelator(i, j, -1, 0, selector)
		f.itelator(i, j, -1, 0, selector)
		f.itelator(i, j, -1, 1, selector)
		f.itelator(i, j, -1, 1, selector)
		f.itelator(i, j, 0, -1, selector)
		f.itelator(i, j, 0, -1, selector)
		f.itelator(i, j, 0, 1, selector)
		f.itelator(i, j, 0, 1, selector)
		f.itelator(i, j, 1, -1, selector)
		f.itelator(i, j, 1, -1, selector)
		f.itelator(i, j, 1, 0, selector)
		f.itelator(i, j, 1, 0, selector)
		f.itelator(i, j, 1, 1, selector)
		f.itelator(i, j, 1, 1, selector)
		f.turn = next
		f.dirty = true
	}
}

func (f *field) String() string {
	str := "  1 2 3 4 5 6 7 8\n"
	black := 0
	white := 0
	f.check()
	for i, line := range f.cells {
		str += strconv.Itoa(i+1) + " "
		for _, val := range line {
			str += val.string()
			switch val.color {
			case BLACK:
				black++
				break
			case WHITE:
				white++
				break
			}
		}
		str += "\n"
	}
	str += "\n" + "[Score] "
	str += new(cell).black().string() + " " + strconv.Itoa(black) + ", "
	str += new(cell).white().string() + " " + strconv.Itoa(white) + "\n"
	return str
}
