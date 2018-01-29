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

	if len(tb.Draw(0, 0)) != 10*10 {
		t.Fatalf("table is printing %v cells, should be printing 100", len(tb.Draw(0, 0)))
	}

	for x, _ := range contents {
		for y, _ := range contents[x] {
			b := tb.OnMouse(x, y, true)
			if b() != "spacer, at"+string(x)+","+string(y) {

				t.Fatalf("mouse returned %s, wanted 'spacer at %v, %v'", b(), x, y)
			}
		}
	}

}
