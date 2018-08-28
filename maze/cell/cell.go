package cell

import (
	. "github.com/nowakf/undefined_behaviour/collection"
)

type Cell Collection

const (
	Tendril Key = Key(Shortval1)
	Body        = Key(Shortval2)
	Food        = Key(Shortval3)
	Wall        = Key(Shortval4)
	Cloud       = Key(Shortval5)
	All         = Key(0xFFFFFFFF)
)

type Key uint64

//you can query however many objects using bitwise OR
//returns a query object which tells you what's there,
//0 if nothign

//returns a string of all the objects in the Cell
func (c Cell) Enumerate() (content string) {
	return Collection(c).String(labels)
}

//Add some objects, or erase them if v is negative
func (c *Cell) Set(k Key, v int) {
	(*Collection)(c).Write(uint64(k), v)
}

func (c Cell) Has(k Key) bool {
	return c.Some(k) == uint64(k)

}
func (c Cell) Some(k Key) uint64 {
	return Collection(c).ReadMultiple(uint64(k))
}
func (c Cell) Read(k Key) int {
	return Collection(c).Read(uint64(k))
}

var labels = map[uint64]string{
	uint64(Tendril): "Tendril",
	uint64(Body):    "Body",
	uint64(Food):    "Food",
	uint64(Wall):    "Wall",
	uint64(Cloud):   "Cloud",
}
