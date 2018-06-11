//object provides two things:
// - an empty object, that contains data and controls access
// - a set of empty structs that have methods for accessing
// the fields of this object
package object

import (
	"sync"
	"ub/common/bitArray"
)

type Object struct {
	m    *sync.Mutex
	data [numCategories]uint32
}

func (r *Object) Get(c category, i index, t dataType) int {

	r.m.Lock()
	defer r.m.Unlock()

	return bitArray.Get32(r.data[c], uint(i), uint(t))

	//i would refer to the offset, t to the length of the object in memory
}

func (r *Object) Set(c category, i index, t dataType, newValue uint8) {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[c] = bitArray.Write32(r.data[c], newValue, uint(i), uint(t))

}

func (r Object) Title(c category, i index, t dataType) string {
	return descriptions[c][i].Title()
}

func (r Object) Blurb(c category, i index, t dataType) string {
	return descriptions[c][i].Blurb()
}

const (
	body category = iota
	mind
	social
	identity
	numCategories //this just tells you how many there are in this block
)

type category int

const (
	boolean  dataType = 1
	nibble            = 4
	bitfield          = 8
)

type dataType int

type index int
