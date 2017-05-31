package game

type cell struct {
	color color
	next  [8]*cell
}

func (c *cell) String() string {
	return c.color.String()
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

func (c *cell) iterator(i int, f func(*cell) bool) {
	for now := c; now != nil && f(now); now = now.next[i] {
	}
}

func (c *cell) check(color color) []*cell {
	var result, cells []*cell
	var status int

	callback := func(now *cell) bool {
		cells = append(cells, now)
		if c == now {
			status = 0
			cells = []*cell{c}
			return c.isnot(BLACK | WHITE)
		} else if status == 0 && now.is((BLACK|WHITE) & ^color) {
			status = 1
			return true
		} else if status == 1 && now.is((BLACK|WHITE) & ^color) {
			return true
		} else if status == 1 && now.is(color) {
			result = append(result, cells...)
			c.change(PUTABLE)
			return false
		} else {
			return false
		}
	}

	c.iterator(0, callback)
	c.iterator(1, callback)
	c.iterator(2, callback)
	c.iterator(3, callback)
	c.iterator(4, callback)
	c.iterator(5, callback)
	c.iterator(6, callback)
	c.iterator(7, callback)

	return result
}
