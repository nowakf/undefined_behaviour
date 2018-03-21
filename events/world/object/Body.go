package object

type Body struct {
}

const (
	strength index = iota * index(nibble)
	intelligence
	willpower
)
const (
	wounds = willpower + iota*index(bitfield)
	feelings
)

func (b *Body) Strength() (category, index, dataType) {
	return body, strength, nibble
}

func (b *Body) Wounds() *wound {
	return &wound{}
}

type wound struct {
}

const (
	head = iota
	leg
)

func (w *wound) Head() (category, index, dataType) {
	return body, wounds + head, boolean
}
