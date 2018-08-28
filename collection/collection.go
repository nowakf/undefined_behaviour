package collection

import "math/bits"
import "fmt"

const (
	Shortval1 uint64 = 1 << iota
	Shortval2
	Shortval3
	Shortval4
	Shortval5
	Shortval6
	Shortval7
	Shortval8
	Longval1
	Longval2
	Longval3
	Longval4
	Bool1
	Bool2
	Bool3
	Bool4
	Bool5
	Bool6
	Bool7
	Bool8
	Bool9
	Bool10
	Bool11
	Bool12
	Bool13
	Bool14
	Bool15
	Bool16
	Bool17
	Bool18
	Bool19
	Bool20
	Bool21
	Bool22
	Bool23
	Bool24
	Bool25
	Bool26
	Bool27
	Bool28
	Bool29
	Bool30
	Bool31
	Bool32
	Bool33
	Bool34
	Bool35
	Bool36
	Bool37
	Bool38
	Bool39
	Bool40
	Bool41
	Bool42
	Bool43
	Bool44
	Bool45
	Bool46
	Bool47
	Bool48
	Bool49
	Bool50
	Bool51
	Bool52
)

type nibblePair uint8

func (n nibblePair) get(part bool) int {
	if part {
		return int(n >> 4)
	} else {
		return int(n & 0x0F)
	}
}
func (n *nibblePair) set(part bool, val uint8) {
	if val > 15 {
		panic("overflows nibble")
	}
	if part {
		*n &^= 0xF0
		*n |= nibblePair(val << 4)
	} else {
		*n &^= 0x0F
		*n += nibblePair(val)

	}
}

type Collection struct {
	bools  uint64
	uint8s [4]uint8
	uint4s [4]nibblePair
}

//Reads a single field - will just read the first if given a larger query - returns -1 if there is no such value
func (c Collection) Read(query uint64) int {
	offset := bits.LeadingZeros64(query)
	if offset < (len(c.uint4s) * 2) {
		return c.uint4s[offset/2].get(offset%2 == 0)
	}
	if offset < (len(c.uint4s)*2 + len(c.uint8s)) {
		return int(c.uint8s[offset-(len(c.uint4s)*2)])
	}

	return int((c.bools&query)>>(64-uint(offset))) - 1
}

//Reads as many fields as it is given. Will not return value of scalars
//- just their presence or non-presence.
func (c Collection) ReadMultiple(query uint64) (intersection uint64) {
	return c.bools & query
}

func (c *Collection) Write(at uint64, val int) {

	if val > 255 {
		panic(fmt.Sprintf("%d is out of bounds", val))
	}

	if val > 0 {
		c.bools |= at
	} else {
		c.bools &^= at
		val = 0
	}

	for i := 0; i < len(c.uint8s); i++ {
		// << by 8 places already, to account for the uint4s
		if 0x00000100<<(64-uint(i))&at == 0 {
			continue
		}
		c.uint8s[i] = uint8(val)
	}

	for i := 0; i < (len(c.uint4s) * 2); i++ {
		if 0x00000001<<(64-uint(i))&at == 0 {
			continue
		}
		c.uint4s[i/2].set(i%2 == 0, uint8(val))
	}

}

func (c Collection) Count() int {
	return bits.OnesCount64(c.bools)
}

func (c Collection) String(labels map[uint64]string) (output string) {
	for k, label := range labels {
		if val := c.Read(k); val > 0 {
			output += fmt.Sprintf("%s, %d, ", label, val)
		}

	}
	return
}
