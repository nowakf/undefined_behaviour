package elements

import c "cthu3/common"
import cp "cthu3/ui/elements/components"

type Table struct {
	contents [][]UiElement
	*cp.Rect
}

func NewTable(h, w int, contents [][]UiElement) *Table {
	t := new(Table)
	t.contents = contents
	t.Rect = cp.NewRect(h, w)
	return t
}

//passes offsets onto contained elements.
func (t *Table) Draw(xoffset, yoffset int) []c.Cell {

	cells := make([]c.Cell, 0)

	width := 0

	for x, array := range t.contents {

		height := 0

		for _, element := range t.contents[x] {

			cells = append(cells, element.Draw(xoffset+width, yoffset+height)...)
			height += element.H()
		}
		width += array[0].W()
	}
	return cells
}

//returns the function associated with the element under the mouse
func (t *Table) OnMouse(x, y int, clicked bool) func() string {

	accumulatedWidths := 0

	for i, column := range t.contents {

		if x <= column[0].W() && x >= accumulatedWidths {
			accumulatedHeights := 0
			for j, box := range column {
				if y < box.H() && y > accumulatedHeights {
					//send the click event to the element
					return box.OnMouse()

				} else {
					accumulatedHeights += box.H()
				}
			}
		} else {
			accumulatedWidths += column[0].W()
		}
	}
	return func() string { return "nothing under cursor at" + string(x) + string(y) }
}