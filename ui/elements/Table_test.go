package elements

import "testing"

var mockTable Table = NewTable(Root(), 100, 100)

var mockArray cumulative = make(cumulative, 0)

func Test_addWidth(t *testing.T) {
}
func Test_insert(t *testing.T) {}
func Test_absWidth(t *testing.T) {
}
func TestWriteToCell(t *testing.T) {
	a := NewSpacer(mockTable.Node, 10, 10)
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10; y++ {
			a = NewTextbox(mockTable.Node, 10, 10, string(x)+string(y))
			mockTable.WriteToCell(x, y, a)
		}
	}
}
func TestResize(t *testing.T) {
}
