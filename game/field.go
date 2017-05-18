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

func (f *field) check(i int, j int, i_offset int, j_offset int) []*cell {
	result := []*cell{}
	cells := []*cell{}
	step := 0
	f.itelator(i, j, i_offset, j_offset, func(c *cell) bool {
		ret := true
		cells = append(cells, c)
		switch step {
		case 0:
			ret = c.isnot(BLACK | WHITE)
			step++
			break
		case 1:
			ret = c.is(BLACK|WHITE) && c.isnot(f.turn)
			step++
			break
		case 2:
			ret = c.is(BLACK|WHITE) && c.isnot(f.turn)
			if c.is(f.turn) {
				result = cells
				f.cells[i][j].change(PUTABLE)
			}
			break
		}
		return ret
	})
	return result
}

func (f *field) checkCell(i int, j int) []*cell {
	cells := append(f.check(i, j, -1, -1), f.check(i, j, -1, 0)...)
	cells = append(cells, f.check(i, j, -1, 1)...)
	cells = append(cells, f.check(i, j, 0, -1)...)
	cells = append(cells, f.check(i, j, 0, 1)...)
	cells = append(cells, f.check(i, j, 1, -1)...)
	cells = append(cells, f.check(i, j, 1, 0)...)
	cells = append(cells, f.check(i, j, 1, 1)...)

	return cells
}

func (f *field) checkCells() int {
	if f.dirty == false {
		return -1
	}
	result := 0
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
			result += len(f.checkCell(i, j))
		}
	}
	return result
}

func (f *field) next() {
	switch f.turn {
	case BLACK:
		f.turn = WHITE
		break
	case WHITE:
		f.turn = BLACK
		break
	}
}

func (f *field) Select(i int, j int) {
	cells := f.checkCell(i, j)
	if f.cells[i][j].is(PUTABLE) {
		f.cells[i][j].change(f.turn)
		for _, cell := range cells {
			cell.change(f.turn)
		}
		f.next()
		f.dirty = true
	}
}

func (f *field) String() string {
	str := "  1 2 3 4 5 6 7 8\n"
	black := 0
	white := 0
	if f.checkCells() == 0 {
		f.next()
		if f.checkCells() == 0 {
			return "Game Over\n"
		}
	}
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
