package elements_test

import "testing"
import e "cthu3/ui/elements"

func TestOnMouse(t *testing.T) {

	contents := make([][]e.UiElement, 10)

	for i, _ := range contents {
		contents[i] = make([]e.UiElement, 10)
		for j, _ := range contents[i] {
			contents[i][j] = e.NewSpacer(1, 1)
		}
	}

	tb := e.NewTable(10, 10, contents)

	for x, _ := range contents {
		for y, _ := range contents[x] {
			b, x1, y1 := tb.OnMouse(x, y, true)
			if b != true || x1 != x || y1 != y {
				println(b, x, y, x1, y1)
				m := string(x1) + "," + string(y1)
				n := string(x) + "," + string(y)

				t.Fatalf("mouseclick returned coordinate (%s) and boolean (%v), wanted (%s) and (%v)", m, b, n, true)
			}
		}
	}

}
