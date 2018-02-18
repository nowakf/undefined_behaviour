package elements

import "testing"

func TestAbsolute(t *Testing.T) {
}

func TestRelative(t *testing.T) {
	d := Dimensions{Length{100, true}, Length{100, true}}
	p := Root()
	s := NewSpacer(p, &d)

	for i := 0; i < 100; i++ {
		r := Relative(i, s.H)
		if r(0) != i {
			t.Fatalf("Relative returns %v, expected %v", r(0), i)
		}

	}

}
func TestResize(t *testing.T) {
	d := Dimensions{Length{100, true}, Length{100, true}}
	p := Root()
	s := NewSpacer(p, &d)

	r := Relative(50, s.H)

	s.Resize(50, 50)

	if r(0) != 25 {
		t.Fatalf("Resize returns %v, should return 25", r(0))
	}
}
