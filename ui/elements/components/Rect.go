package components

type Rect struct {
	h, w int
}

func NewRect(h, w int) *Rect {
	r := new(Rect)
	r.h, r.w = h, w
	return r
}
func (r *Rect) H() int { return r.h }

func (r *Rect) W() int { return r.w }
