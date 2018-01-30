package elements

import c "cthu3/common"
import "fmt"

type Table struct {
	contents [][]UiElement
	*rect
	xorigin, yorigin int
}

func NewTable(h, w int) *Table {
	t := new(Table)
	t.contents = make([][]UiElement, 0)
	t.rect = newrect(h, w)
	return t
}
func (t *Table) WriteToCell(x, y int, e UiElement) {

	//you fuck you

	if x >= t.ColumnCount() || y >= t.RowCount(x) {
		t.AddCells(x+1-t.ColumnCount(), y+1-t.RowCount(x))
	}

	t.contents[x][y] = e

	if e.W() > t.contents[x][0].W() {
		t.contents[x][0].SetW(e.W())
	}

}

func (t *Table) AddCells(deltax, deltay int) {

	n := make([][]UiElement, t.ColumnCount()+deltax)

	for x := 0; x < len(n); x++ {

		n[x] = make([]UiElement, t.RowCount(x)+deltay)

		for y := 0; y < len(n[x]); y++ {

			//you know what, fuck it
			if x < len(t.contents) && y < len(t.contents[x]) && t.contents[x][y] != nil {
				n[x][y] = t.contents[x][y]
			} else {
				n[x][y] = NewSpacer(1, 1)
			}
		}

	}

	t.contents = n
}

func (t *Table) ColumnCount() int {
	return len(t.contents)
}

func (t *Table) RowCount(row int) int {

	if len(t.contents) > row {
		return len(t.contents[row])
	} else {
		return 0
	}
}

//passes offsets onto contained elements.
func (t *Table) Draw(xoffset, yoffset int) []c.Cell {

	//I kinda prefer initializing this here, because otherwise I think you'd have to manually specify an origin
	t.xorigin, t.yorigin = xoffset, yoffset

	cells := make([]c.Cell, 0)

	width := 0

	for x, column := range t.contents {

		height := 0

		for _, element := range t.contents[x] {

			cells = append(cells, element.Draw(xoffset+width, yoffset+height)...)

			height += element.H()
		}

		width += column[0].W()
	}
	return cells
}

//returns the function associated with the element under the mouse
func (t *Table) OnMouse(x, y int, pressed bool, released bool) func() string {

	accumulatedWidths := t.xorigin

	for _, column := range t.contents {

		if x < accumulatedWidths+column[0].W() && x >= accumulatedWidths {

			accumulatedHeights := t.yorigin

			for _, box := range column {
				if y < accumulatedHeights+box.H() && y >= accumulatedHeights {
					//send the click event to the element
					return box.OnMouse(x, y, pressed, released)

				} else {
					accumulatedHeights += box.H()
				}
			}
		} else {
			accumulatedWidths += column[0].W()
		}
	}
	return func() string { return fmt.Sprintf("mouse probably out of bounds at %v,%v", x, y) }
}
