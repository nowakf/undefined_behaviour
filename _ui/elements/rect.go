package elements

import (
	"github.com/faiface/pixel"
	c "ub/common"
)

type rect struct {
	h, w       int
	foreground pixel.RGBA
	background pixel.RGBA
	//
}

func newRect(h, w int) rect {
	r := new(rect)
	r.foreground = c.Red
	r.background = c.Red
	r.h, r.w = h, w
	return *r
}

func (r *rect) H() int {
	return r.h
}
func (r *rect) W() int {
	return r.w
}

func (r *rect) Resize(h int, w int) {
	r.h, r.w = h, w
}
