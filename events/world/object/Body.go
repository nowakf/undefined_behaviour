package object

type Body struct {
	stats nibbleArray
}

const (
	STR int = iota * 4
)

func (b *Body) Get() *nibbleArray {
	return &b.stats
}

type index int
