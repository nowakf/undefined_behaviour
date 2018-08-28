//package object is a reasonably safe uint64 block that can
//contain various values
package object

import (
	"fmt"
	"math/bits"
)

var Keys = [...]Key{
	BOOL0,
	BOOL1,
	BOOL2,
	BOOL3,
	BOOL4,
	BOOL5,
	BOOL6,
	BOOL7,
	BOOL8,
	BOOL9,
	BOOL10,
	BOOL11,
	BOOL12,
	BOOL13,
	BOOL14,
	BOOL15,
	NIBBLE16,
	NIBBLE20,
	NIBBLE24,
	NIBBLE28,
	NIBBLE32,
	NIBBLE36,
	NIBBLE40,
	NIBBLE44,
	UINT8_48,
	UINT8_54,
}

const (
	FIRST_BOOL   = 0
	FIRST_NIBBLE = 16
	FIRST_UINT8  = 25
)

const (
	BOOL0 Key = iota
	BOOL1
	BOOL2
	BOOL3
	BOOL4
	BOOL5
	BOOL6
	BOOL7
	BOOL8
	BOOL9
	BOOL10
	BOOL11
	BOOL12
	BOOL13
	BOOL14
	BOOL15
	lastBoolBoundary
)
const (
	NIBBLE16 Key = lastBoolBoundary + iota*4
	NIBBLE20
	NIBBLE24
	NIBBLE28
	NIBBLE32
	NIBBLE36
	NIBBLE40
	NIBBLE44
	lastNibbleBoundary
)
const (
	UINT8_48 Key = lastNibbleBoundary + iota*8
	UINT8_54
)

const (
	MAX      = 0xFFFFFFFFFFFFFFFF
	boolMask = MAX >> (64 - lastBoolBoundary)
)
const (
	_bool   = 1
	_nibble = 4
	_uint8  = 8
)

type Object uint64

type Key uint

func (k Key) Length() uint {
	switch {
	case k < lastBoolBoundary:
		//it's a bool
		return _bool
	case k < lastNibbleBoundary:
		//its a nible
		return _nibble
	default:
		//it's a uint8
		return _uint8
	}
}
func (k Key) checkFit(newVal int) int {
	switch k.Length() {
	case _bool:
		if newVal < 2 && newVal >= 0 {
			return newVal
		}
	case _nibble:
		if newVal < 16 && newVal >= 0 {
			return newVal
		}
	case _uint8:
		if newVal < 256 && newVal >= 0 {
			return newVal
		}
	}
	fmt.Printf("attempted to set out of bounds value %d, for Key %d, \n", newVal, k)

	return 0
}

//get an object using one of the predefined constant Keys
func (o Object) Get(k Key) int {
	o <<= 64 - uint(k) - k.Length()
	o >>= 64 - k.Length()
	return int(o)
}

//set an object using one of the predefined constant Keys
func (o *Object) Set(k Key, newVal int) *Object {
	//if it's out of bounds, set to zero and log error
	newVal = k.checkFit(newVal)

	var mask uint64 = MAX

	mask <<= 64 - k.Length()
	//size the mask

	mask >>= 64 - uint(k) - k.Length()
	//move it to the right location

	*o &^= Object(mask)
	//clear the area

	newVal <<= k
	//move the new value to the same offset

	*o |= Object(newVal)

	return o
}

//checks if the object 'this' is a subset of 'o'
//would like to do this with bitwise techniques -
//need to do research
func (o Object) Superset(this Object) bool {

	if (this&boolMask)&^(o&boolMask) != 0 {
		return false
	}

	for i := FIRST_NIBBLE; i < FIRST_UINT8; i++ {
		if o.Get(Keys[i]) < this.Get(Keys[i]) {
			return false
		}
	}
	for i := FIRST_UINT8; i < len(Keys); i++ {
		if o.Get(Keys[i]) < this.Get(Keys[i]) {
			return false
		}

	}
	return true

}

//returns fields that have a value -
//can be optimized a lot if necessary
func (o Object) Fields() (filled []Key) {
	for _, key := range Keys {
		if o.Get(key) != 0 {
			filled = append(filled, key)
		}
	}
	return
}

//debug only
func (o Object) PrintBinary() string {
	s := ""
	var rightmostBit uint64
	rightmostBit = 0x0000000000000001

	for i := 0; i <= 64; i++ {
		o1 := o

		o1 <<= 64 - uint(bits.Len64(uint64(o1)))
		if rightmostBit&(uint64(o1)>>uint(i)) == 1 {
			s = "█" + s
		} else {
			s = " " + s
		}

	}
	return s
}

func (o Object) RelativeComplement(this Object) (diff []Key) {

	for _, key := range Keys {
		if o.Get(key) < this.Get(key) {
			diff = append(diff, key)
		}
	}
	return diff
}
