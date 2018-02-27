package elements

import (
	"math"
	"math/rand"
)

type container struct {
	id int
	*Node
	*rect
	rW, rH float64
}

func NewContainer(this UiElement, parent *Node, h, w int) *container {
	c := new(container)
	if parent == nil || parent.element == nil || parent.element.H() == 0 || parent.element.W() == 0 {
		c.rW, c.rH = 1, 1
	} else {
		//this reccords the relation of parent to child
		c.rH = float64(h) / float64(parent.element.H())
		c.rW = float64(w) / float64(parent.element.W())
	}

	r := newRect(h, w)

	c.rect = &r

	c.Node = newNode(parent, this)

	c.id = rand.Int()

	return c

}

func (c *container) GetRatio() (float64, float64) {
	return c.rH, c.rW
}

func (c *container) Resize(h, w int) {
	c.rect.Resize(h, w)
	for _, child := range c.GetChildren() {
		hr, wr := child.GetElement().GetRatio()
		fh := float64(h) * hr
		fw := float64(w) * wr
		resizedH := int(math.Ceil(fh))
		resizedW := int(math.Ceil(fw))
		child.GetElement().Resize(resizedH, resizedW)
	}
}
