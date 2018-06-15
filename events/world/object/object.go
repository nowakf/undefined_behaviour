package object

import "fmt"

const (
	clearMask = 0xFFFFFFFFFFFFFFFF
)
const (
	MemberOfCult key = iota
	MemberOfAntiCult
	LawAbiding
	//wounds
	MissingLimb
	MissingEye
	HeartProblems
	Schizophrenia
	Paranoia
	Agoraphobia
	//
	Depression
	Plague
	_
	_
	_
	_
	_
	lastBoolBoundary
)
const (
	_ key = lastBoolBoundary + iota*4
	STR
	CHA
	WIS
	INT
	WIL
	SAN
	_
	lastNibbleBoundary
)
const (
	ID key = lastNibbleBoundary + iota*8
	STRESS
)

const (
	_bool   = 1
	_nibble = 4
	_uint8  = 8
)

type Object uint64

type key uint

func (o Object) blockLen(k key) (length uint) {
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
func (o Object) checkFit(k key, newVal int) int {
	switch o.blockLen(k) {
	case _bool:
		if newVal < 2 && newVal > 0 {
			return newVal
		}
	case _nibble:
		if newVal < 16 && newVal > 0 {
			return newVal
		}
	case _uint8:
		if newVal < 256 && newVal > 0 {
			return newVal
		}
	}
	fmt.Printf("attempted to set out of bounds value %d, for key %d", newVal, k)

	return 0
}

//get an Object using one of the predefined constant keys
func (o Object) Get(k key) int {
	o <<= 64 - uint(k) - o.blockLen(k)
	o >>= 64 - o.blockLen(k)
	return int(o)
}

//set an Object using one of the predefined constant keys
func (o *Object) Set(k key, newVal int) *Object {
	//if it's out of bounds, set to zero and log error
	newVal = o.checkFit(k, newVal)

	var mask uint64 = clearMask

	mask <<= 64 - o.blockLen(k)
	//size the mask

	mask >>= 64 - uint(k) - o.blockLen(k)
	//move it to the right location

	*o &^= Object(mask)
	//clear the area

	newVal <<= k
	//move the new value to the same offset

	*o |= Object(newVal)

	return o
}
