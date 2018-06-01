package elements

import "testing"
import "math/rand"

var mockTable *Table = NewTable(Root(), 100, 100)

var mockArray cumulative = make(cumulative, 0)

func Test_addWidth(t *testing.T) {
	length := 100
	widths := make([]int, length)
	for i := 0; i < length; i++ {
		widths[i] = rand.Int()
		mockArray.addWidth(i, widths[i])
	}
	//test summing widths:
	sum := 0
	for _, v := range widths {
		sum += v
	}
	last := mockArray[len(mockArray)-1]
	if sum != last {
		t.Errorf("got a total width of %v, expected %v", last, sum)
	}

}

func Test_absWidth(t *testing.T) {
	mockArray = make(cumulative, 0)
	index := make([]int, 101)
	width := make([]int, 101)
	for i := 0; i < 100; i++ {
		index[i] = rand.Intn(1000)
		width[i] = rand.Intn(1000)
		mockArray.addWidth(index[i], width[i])
	}
	for i := 0; i < len(width); i++ {
		mIndex := index[i]
		got := mockArray.absWidth(mIndex)
		expected := width[i]
		if got != expected {
			t.Errorf("absolute width returned %v, expected %v,", got, expected)
			t.Errorf("the index was %v,", index[i])

		}
	}
}
