package elements

type rect struct {
	h, w int
}

func newrect(h, w int) *rect {
	r := new(rect)
	r.h, r.w = h, w
	return r
}
func (r *rect) H() int { return r.h }

func (r *rect) W() int { return r.w }

func (r *rect) SetW(w int) {
	r.w = w
}

func (r *rect) SetH(h int) {
	r.h = h
}
