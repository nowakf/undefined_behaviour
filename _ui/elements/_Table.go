package elements

import c "ub/common"
import "fmt"

type Table struct {
	*container
	columns          []block
	rows             []block
	xorigin, yorigin int
}

func NewTable(parent *Node, h, w int) *Table {
	t := new(Table)
	t.container = NewContainer(t, parent, h, w)
	t.rows = make([]block, 0)
	t.columns = make([]block, 0)
	return t
}

func (t *Table) WriteToCell(x, y int, e UiElement) {

	oldXLength := len(t.columns)
	oldYLength := len(t.rows)

	if x >= len(t.columns) {
		t.columns = t.extendBlocks(x+1, t.columns, len(t.rows))
	}
	if y >= len(t.rows) {
		t.rows = t.extendBlocks(y+1, t.rows, len(t.columns))
	}

	if y >= len(t.columns[x].cont) {
		t.columns = t.extendContents(y+1, oldYLength, t.columns)
	}
	if x >= len(t.rows[y].cont) {
		t.rows = t.extendContents(x+1, oldXLength, t.rows)
	}

	if e.W() > t.columns[x].breadth {
		t.columns[x].breadth = e.W()
	}
	if e.H() > t.rows[y].breadth {
		t.rows[y].breadth = e.H()
	}

	t.columns[x].cont[y] = e
	t.rows[y].cont[x] = e

}

//so len(t.collumns) = len(t.row[x])
//extend block increases the length of its input row or column, and adds cells
func (t *Table) extendBlocks(newLength int, a []block, otherAxis int) []block {

	l := len(a)
	//add to the first axis
	a = append(a, make([]block, newLength+1-len(a))...)
	//new arrays, of the same length as the other axis
	for i := l; i < len(a); i++ {
		a[i] = block{breadth: 0, cont: make([]UiElement, otherAxis)}
	}

	return a
}

func (t *Table) extendContents(newLength int, oldLength int, a []block) []block {
	//make the extension for the other axis
	extra := make([]UiElement, newLength+1-oldLength)
	for i, _ := range a {
		a[i].cont = append(a[i].cont, extra...)
	}
	return a

}

func (t *Table) Resize(h, w int) {
	//this can absolutely be improved if necessary
	for _, column := range t.columns {
		maxW, _ := column.maxHW()
		column.breadth = maxW
	}
	for _, row := range t.rows {
		_, maxH := row.maxHW()
		row.breadth = maxH
	}
	t.container.Resize(h, w)
}

//passes offsets onto contained elements.
func (t *Table) Draw(xoffset, yoffset int) []c.Cell {

	//I kinda prefer initializing this here, because otherwise I think you'd have to manually specify an origin
	t.xorigin, t.yorigin = xoffset, yoffset

	cells := make([]c.Cell, 0)

	width := 0

	for _, column := range t.columns {

		height := 0

		for i, row := range t.rows {

			if column.cont[i] != nil {
				cells = append(cells, column.cont[i].Draw(xoffset+width, yoffset+height)...)
			}

			height += row.breadth

		}

		width += column.breadth
	}
	return cells
}

func (t *Table) Identify() string {
	return fmt.Sprintf("table, origin (%v, %v,) w of: %v, h of: %v", t.xorigin, t.yorigin, t.w, t.h)
}

func (t *Table) GetLast(x, y int) UiElement {

	accumulatedWidths := t.xorigin

	for i, column := range t.columns {

		if x < accumulatedWidths+column.breadth && x >= accumulatedWidths {

			accumulatedHeights := t.yorigin

			for _, row := range t.rows {
				if y < accumulatedHeights+row.breadth && y >= accumulatedHeights {
					if row.cont[i] != nil {
						return row.cont[i].GetLast(x, y)
					} else {
						return nil
					}

				} else {
					accumulatedHeights += row.breadth
				}
			}
		} else {
			accumulatedWidths += column.breadth
		}
	}
	return nil
}

type block struct {
	breadth int
	cont    []UiElement
}

func (b *block) maxHW() (int, int) {
	maxH := 0
	maxW := 0
	for _, element := range b.cont {
		if element != nil {
			if maxH < element.H() {
				maxH = element.H()
			}
			if maxW < element.W() {
				maxW = element.W()
			}
		}
	}
	return maxH, maxW
}
