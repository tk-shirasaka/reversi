package game

import (
	"strconv"
)

type field struct {
	turn  color
	dirty bool
	cells [8][8]*cell
}

func Init() *field {
	f := new(field)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			f.cells[i][j] = new(cell)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i > 0 && j > 0 {
				f.cells[i][j].next[0] = f.cells[i-1][j-1]
			}
			if i > 0 {
				f.cells[i][j].next[1] = f.cells[i-1][j]
			}
			if i > 0 && j < 7 {
				f.cells[i][j].next[2] = f.cells[i-1][j+1]
			}
			if j > 0 {
				f.cells[i][j].next[3] = f.cells[i][j-1]
			}
			if j < 7 {
				f.cells[i][j].next[4] = f.cells[i][j+1]
			}
			if i < 7 && j > 0 {
				f.cells[i][j].next[5] = f.cells[i+1][j-1]
			}
			if i < 7 {
				f.cells[i][j].next[6] = f.cells[i+1][j]
			}
			if i < 7 && j < 7 {
				f.cells[i][j].next[7] = f.cells[i+1][j+1]
			}
		}
	}

	f.turn = BLACK
	f.dirty = true
	f.cells[3][3].change(BLACK)
	f.cells[3][4].change(WHITE)
	f.cells[4][3].change(WHITE)
	f.cells[4][4].change(BLACK)

	return f
}

func (f *field) checkCells() int {
	if f.dirty == false {
		return -1
	}
	result := 0
	for _, line := range f.cells {
		for _, c := range line {
			if c.isnot(BLACK | WHITE) {
				c.change(BLANK)
			}
		}
	}
	for _, line := range f.cells {
		for _, c := range line {
			result += len(c.check(f.turn))
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
	cells := f.cells[i][j].check(f.turn)
	if f.cells[i][j].is(PUTABLE) {
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
			str += val.String()
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
	str += new(cell).change(BLACK).String() + " " + strconv.Itoa(black) + ", "
	str += new(cell).change(WHITE).String() + " " + strconv.Itoa(white) + "\n"
	str += "[Next]  " + new(cell).change(f.turn).String()
	return str
}
