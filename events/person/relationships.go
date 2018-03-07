package person

import "sort"

//relationships encompass all kinds of relationship - including to the world, and to organizations...
//
//it implements sort.Interface, so yeah,
type contacts struct {
	loved uint8
	hated uint8
	*relationships
}

func (c *contacts) Update() {
	sort.Sort(c.relationships)
	c.loved = c.relationships[len(*c.relationships)-1].entity
	c.hated = c.relationships[0].entity
}

type relationships [256]relationship

func (r *relationships) Len() int {
	return len(*r)
}
func (r *relationships) Less(i, j int) bool {
	return (*r)[i].state <= (*r)[j].state
}
func (r *relationships) Swap(i, j int) {
	temp := (*r)[i]
	(*r)[i] = (*r)[j]
	(*r)[j] = temp
}

type relationship struct {
	entity       uint8
	state        int
	descriptions []string
	deltas       []int
}

//prolly need some kind of lock on this
func (r *relationship) Event(delta int, description string, actor int) {

}
