package person

import "sync"

//stats is a compressed representation of the stats of an entity. As such, it isn't particularly human-readable
type stats struct {
	m         sync.Mutex
	nibbles   uint64
	boolean8s uint64
	magic4s   uint64
}

const (
	madness nibble_Index = iota * 4
	stress
	membership
	alignment
	intelligence
	willpower
	strength
	constitution
	charisma
	health
	//
	//
	//
	// 64
)

type nibble_Index uint

//probably relocate this to some kind of package e.g
//skills?
const (
	traits boolean8_index = iota * 8
	motives
	wounds
	normal
	criminal
	difficult
	social
	// 64
)

type boolean8_index uint

//probably relocate this to some kind of package
const (
	sober magic4_index = iota * 8
	kolmogorov
	frechet
	urysohn
	r_hausdorff
	tychonoff
	n_hausdorff
	borgia
)

type magic4_index uint

//Get returns the complete uint4-layout object
//a uint4 layout is one where each 4 bits refers
//to a value between 0 and 15 (e.g one hex digit)
func (s *stats) GetNibbles() uint64 {
	return s.nibbles
}

//Get returns the complete boolean-layout object
//A boolean layout is where each bit refers to
//a boolean value
func (s *stats) GetBooleans() uint64 {
	return s.boolean8s
}

//Get returns the complete magic-layout object
//a magic layout is where each eight bits refers to
//a school, as defined in the magic4_index const
//and each 4 bits within that eight bits refers to a
//spell within that school
func (s *stats) GetMagic() uint64 {
	return s.magic4s
}
func (s *stats) Madness() *nibble {
	return &nibble{&s.m, s.nibbles, madness}
}
func (s *stats) Motives() *boolean8 {
	return &boolean8{&s.m, s.boolean8s, motives}
}

type nibble struct {
	m *sync.Mutex
	uint64
	nibble_Index
}

func (u *nibble) Read() int {
	u.m.Lock()
	defer u.m.Unlock()
	return int(u.uint64 >> u.nibble_Index)
}

type boolean8 struct {
	m *sync.Mutex
	uint64
	boolean8_index
}

func (b *boolean8) Read(index int) bool {
	b.m.Lock()
	defer b.m.Unlock()
	bitarray := uint8(b.uint64 >> b.boolean8_index)
	_ = bitarray
	//check the bit
	return true
}

type magic4 struct {
	m *sync.Mutex
	uint64
	magic4_index
}
