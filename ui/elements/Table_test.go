package elements

import "testing"
import "fmt"

func TestNewTable(t *testing.T) {
	width, height := 20, 20
	ta := NewTable(height, width)
	if ta.H() != height || ta.W() != width {
		t.Fatalf("table is the wrong size! expected %vx%v, got %vx%v", height, width, ta.H(), ta.W())
	}
}

func TestAddCells(t *testing.T) {
	ta := NewTable(40, 40)
	expectC, expectR := 3, 3
	ta.AddCells(expectC, expectR)
	if len(ta.contents) != expectC {
		t.Fatalf("1. wrong column no, expected %v, got %v", expectC, len(ta.contents))
	}
	for _, column := range ta.contents {
		if len(column) != expectR {
			t.Fatalf("1. wrong row no, expected %v, got %v", expectR, len(column))
		}
	}
	column := []UiElement{
		NewSpacer(3, 3),
		NewSpacer(3, 3),
		NewSpacer(3, 3),
	}
	ta.contents = [][]UiElement{column, column, column}

	expectC, expectR = 3, 3
	if len(ta.contents) != expectC {
		t.Fatalf("2. wrong column no, expected %v, got %v", expectC, len(ta.contents))
	}

	for _, column := range ta.contents {
		if len(column) != expectR {
			t.Fatalf("wrong row no, expected %v, got %v", expectR, len(column))
		}
	}

	if ta.ColumnCount() != expectC {
		t.Fatalf("ColumnCount() reported %v rows, we wanted %v!", ta.ColumnCount(), expectC)
	}

	//check if the RowCounter is working:

	expectC, expectR = 6, 6

	ta.AddCells(3, 3)
	if len(ta.contents) != expectC {
		t.Fatalf("3.wrong column no, expected %v, got %v", expectC, len(ta.contents))
	}
	for i, column := range ta.contents {
		if len(column) != expectR {
			t.Fatalf("3 .wrong row no, expected %v, got %v, at column %v", expectR, len(column), i)
		}
	}

}

func TestWriteToCell(t *testing.T) {
	ta := NewTable(5, 5)
	for i := 0; i < 5; i++ {
		s := fmt.Sprintf("test box %v", i)
		ta.WriteToCell(0, i, NewTextButton(5, s, func() string {
			return s
		}))
	}
	for x := 0; x < len(ta.contents); x++ {
		for y := 0; y < len(ta.contents[x]); y++ {
			got := ta.contents[x][y].OnMouse(x, y, false, true)()
			expected := fmt.Sprintf("test box %v", y)
			if got != expected {
				t.Fatalf("you got %v, and expected %v", got, expected)
			}

		}
	}
}

func TestOnMouse(t *testing.T) {
}
