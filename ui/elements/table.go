package elements

import c "ub/common"

type Table struct {
	*container
	origin   coord
	widths   cumulative
	heights  cumulative
	contents map[coord]UiElement
}

func NewTable(parent *Node, h, w int) *Table {
	t := new(Table)
	t.container = NewContainer(t, parent, h, w)
	t.widths = make(cumulative, 0)
	t.heights = make(cumulative, 0)
	t.contents = make(map[coord]UiElement)
	return t
}

func (t *Table) WriteToCell(x, y int, e UiElement) {
	t.widths = t.widths.addWidth(x, e.W())
	t.heights = t.heights.addWidth(y, e.H())
	t.contents[coord{x, y}] = e
}

//resize calls the normal resize, but also adjusts the cumulative widths
func (t *Table) Resize(h, w int) {
	t.container.Resize(h, w)
	t.widths = make(cumulative, 0)
	t.heights = make(cumulative, 0)
	for coord, element := range t.contents {
		t.widths = t.widths.addWidth(coord.x, element.W())
		t.heights = t.heights.addWidth(coord.y, element.H())
	}
}

//change this to just redraw changed - if necessary
//Draws each element in the table at its offset
func (t *Table) Draw(xoffset, yoffset int) []c.Cell {
	t.origin = coord{xoffset, yoffset}
	cells := make([]c.Cell, 0)
	for coord, element := range t.contents {
		cells = append(cells,
			element.Draw(
				t.origin.x+t.widths.origin(coord.x),
				t.origin.y+t.heights.origin(coord.y))...)
	}
	return cells
}

//calls 'GetLast' on the object under the cursor -
//so in a nested table, you still get the last object contained.
func (t *Table) GetLast(x, y int) UiElement {
	for xc, rightBound := range t.widths {
		if x >= t.widths.origin(xc) && x < rightBound {
			for yc, lowerBound := range t.heights {
				if y < lowerBound && y >= t.heights.origin(yc)+t.origin.y {
					element, ok := t.contents[coord{xc, yc}]
					if ok {
						return element.GetLast(x, y)
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}

//id is not actually unique, but it's probably unique enough
func (t *Table) Identify() string {
	return "table" + string(t.id)
}

//I should probably make this one general
type coord struct {
	x, y int
}

//cumulative is a type that stores the cumulative widths of table elements,
type cumulative []int

// 1 1 2 3 4 5
func (c *cumulative) origin(i int) int {
	return (*c)[i] - c.absWidth(i)
}
func (c *cumulative) addWidth(i int, width int) []int {
	initalLen := len(*c)
	if i >= initalLen {
		tail := make([]int, 1+i-initalLen)
		for index, _ := range tail {
			if initalLen > 0 {
				tail[index] = (*c)[initalLen-1]
			}

		}
		*c = append(*c, tail...)
	}
	//fill in the gaps with appropriate values
	c.insert(i, width)
	//add in the value in question
	return *c

}
func (c *cumulative) absWidth(i int) int {
	acc := (*c)[i]
	if i > 0 {
		return acc - (*c)[i-1]
	} else {
		return acc
	}
}
func (c *cumulative) insert(i int, width int) []int {
	orig := c.absWidth(i)
	for i < len(*c) {
		(*c)[i] += width - orig
		i++
	}
	return *c
}

// 0 0
// 0 0 0 0 5
//