package maze

import "image/color"

type arm struct {
	segments vlist
	path     vlist
	state    astate
}
type astate int

const (
	extending astate = iota
	retracting
)

func (a *arm) Print(vs ...vlist) string {
	s := "SEGMENTS:"
	s += a.segments.Print()
	s += "\n PATH:"
	s += a.path.Print()
	return s
}
func (a *arm) assign(p vlist) {
	a.state = retracting
	a.path = p
}

func (a *arm) move(newLoc vector) {
	a.segments.PushFront(newLoc)
	a.path.PushFront(newLoc)
}

func (a *arm) updateLength() {
	switch a.state {
	case extending:
		if a.path.Len() > a.segments.Len() {
			a.segments.PushBack(*a.path.I(a.segments.Len()))
		}
		if a.path.Len() <= 0 {
			a.state = retracting
		}
	case retracting:
		if a.segments.Len() > 0 {
			a.segments.PopBack()
		} else {
			a.state = extending
		}
	}
}

func (a *arm) draw(m *maze) {
	for i := 0; i < a.segments.Len(); i++ {
		m.Set(*a.segments.I(i), Tendril, 0)
	}
	//	for i := 0; i < a.path.Len(); i++ {
	//		m.Set(*a.path.I(i), tendril)
	//	}

}

const (
	none = iota
	N
	S
	NS
	W
	NE
	SE
	NSE
	E
	NW
	SW
	NSW
	EW
	NEW
	SEW
	NSEW
)

type tendril struct {
	name  string
	sigil rune
	color color.RGBA
}

var Tendrils = [...]tendril{
	none: tendril{},
	N: tendril{
		name:  "N",
		sigil: '\u2575',
	},
	S: tendril{
		name:  "S",
		sigil: '\u2577',
	},
	NS: tendril{
		name:  "NS",
		sigil: '\u2502',
	},
	W: tendril{
		name:  "W",
		sigil: '\u2576', //
	},
	NE: tendril{
		name:  "NE",
		sigil: '\u2570', //
	},
	SE: tendril{
		name:  "SE",
		sigil: '\u256D', //
	},
	NSE: tendril{
		name:  "NSE",
		sigil: '\u251C', //
	},
	E: tendril{
		name:  "E",
		sigil: '\u2574', //
	},
	NW: tendril{
		name:  "NW",
		sigil: '\u256F', //
	},
	SW: tendril{
		name:  "SW",
		sigil: '\u256E', //
	},
	NSW: tendril{
		name:  "NSW",
		sigil: '\u2524', //
	},
	EW: tendril{
		name:  "EW",
		sigil: '\u2500',
	},
	NEW: tendril{
		name:  "NEW",
		sigil: '\u2534', //
	},
	SEW: tendril{
		name:  "SEW",
		sigil: '\u252C',
	},
	NSEW: tendril{
		name:  "NSEW",
		sigil: '\u253C',
	},
}
