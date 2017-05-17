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
	f.cells[3][3].change(BLACK)
	f.cells[3][4].change(WHITE)
	f.cells[4][3].change(WHITE)
	f.cells[4][4].change(BLACK)

	return f
}

func (f *field) itelator(i int, j int, i_offset int, j_offset int, callback func(*cell) bool) {
	if i >= 0 && i <= 7 && j >= 0 && j <= 7 && callback(&f.cells[i][j]) {
		f.itelator(i+i_offset, j+j_offset, i_offset, j_offset, callback)
	}
}

func (f *field) check() {
	if f.dirty == false {
		return
	}
	for i, line := range f.cells {
		for j, _ := range line {
			c := &f.cells[i][j]
			if c.isnot(BLACK | WHITE) {
				c.change(BLANK)
			}
		}
	}
	for i, line := range f.cells {
		for j, _ := range line {
			var step int
			checker := func(c *cell) bool {
				ret := true
				if c == &f.cells[i][j] {
					step = 0
				}
				switch step {
				case 0:
					ret = c.is(BLANK)
					break
				case 1:
					ret = c.is(BLACK|WHITE) && c.isnot(f.turn)
					break
				case 2:
					ret = c.is(BLACK|WHITE) && c.isnot(f.turn)
					if c.is(f.turn) {
						f.cells[i][j].change(PUTABLE)
					}
					break
				}
				if ret {
					step++
				}
				return ret
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
	if f.cells[i][j].is(PUTABLE) {
		step := 0
		next := WHITE
		if f.turn == next {
			next = BLACK
		}
		selector := func(c *cell) bool {
			if step == 0 {
				c.color = f.turn
				step = 1
			} else if step == 1 && c.is(next) {
				step = 2
			} else if step == 2 && c.is(f.turn) {
				step = 3
				return false
			} else if step == 3 {
				step = 4
			} else if step == 4 && c.is(next) {
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
	str += new(cell).change(BLACK).string() + " " + strconv.Itoa(black) + ", "
	str += new(cell).change(WHITE).string() + " " + strconv.Itoa(white) + "\n"
	str += "[Next]  " + new(cell).change(f.turn).string()
	return str
}
