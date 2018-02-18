package elements

import "testing"

func TestNewTable(t *testing.T) {
}

func Test_extendBlocks(t *testing.T) {
	n := Root()
	hi := Length{L: 5, Abs: true}
	wi := Length{L: 5, Abs: true}
	di := Dimensions{hi, wi}
	ta := NewTable(n, &di)

	a := make([]block, 0)
	b := make([]block, 0)

	length := 5

	oldLength := len(a)

	a = ta.extendBlocks(length, a, len(b))
	b = ta.extendBlocks(length, b, len(a))
	a = ta.extendContents(length, oldLength, a)

	check := func(name string, x []block, expected int) {
		if len(x) != expected {
			t.Errorf("the block array %s is len(%v), it should be len(%v)", name, len(x), expected)
		}
		for _, y := range x {
			if len(y.cont) != expected {
				t.Errorf("the element array %s is len(%v), it should be len(%v)", name, len(y.cont), expected)
			}

		}
	}
	check("a", a, length-1)
	check("b", b, length-1)

}

func TestWriteToCell(t *testing.T) {
}
