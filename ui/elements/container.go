package elements

import "math"

type container struct {
	*Node
	*rect
	rW, rH float64
}
type Dimensions struct {
	H Length
	W Length
}
type Length struct {
	L   int
	Abs bool //false if relative, true if Absolute
}

func NewContainer(this UiElement, parent *Node, h, w int) *container {
	c := new(container)
	if parent == nil || parent.element == nil || parent.element.H() == 0 || parent.element.W() == 0 {
		c.rW, c.rH = 1, 1
	} else {
		//this reccords the relation of parent to child
		println(parent.element.W(), w, "p.e.W(), width")
		c.rW = float64(w) / float64(parent.element.W())
		println(c.rH, "c.rH")
		c.rH = float64(h) / float64(parent.element.H())

		println(c.rW, "c.rW")
	}

	r := newRect(h, w)
	c.rect = &r

	c.Node = newNode(parent, this)
	return c

}

func (c *container) GetRatio() (float64, float64) {
	return c.rH, c.rW
}

func (c *container) Resize(h, w int) {
	c.rect.Resize(h, w)
	for _, node := range c.GetChildren() {
		hr, wr := node.element.GetRatio()
		fh := float64(h) * hr
		fw := float64(w) * wr
		h = int(math.Ceil(fh)) //so the result is always at least 1
		w = int(math.Ceil(fw))
		node.element.Resize(h, w)
	}
}
