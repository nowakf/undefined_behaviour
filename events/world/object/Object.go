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
	data [CATEGORIES]uint32
}

func (r *Object) Get(c category, i index, t dataType) int {

	r.m.Lock()
	defer r.m.Unlock()

	return bitArray.Get32(r.data[c], uint(i), uint(t))

	//i would refer to the left bound, t to the right
}

func (r *Object) Set(c category, i index, t dataType, newValue uint8) {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[c] = bitArray.Write32(r.data[c], newValue, uint(i), uint(t))

}

func (r *Object) String(c category, i index, t dataType) string {
	return objects[c][i].String()
}

func (r *Object) Describe(c category, i index, t dataType) string {
	return objects[c][i].Describe()
}

const (
	body category = iota
	mind
	social
	identity
	//etc
)

const CATEGORIES = 3

type category int

const (
	boolean  dataType = 1
	nibble            = 4
	bitfield          = 8
)

type dataType int

type index int
