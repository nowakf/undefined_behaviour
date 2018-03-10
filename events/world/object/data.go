package object

type Value interface {
	Get(index int) (value int)
	Set(index int, nu int)
}

type bitfield uint32

type nibbleArray uint32

func (n *nibbleArray) Get(i int) int {
	return 3
}
func (n *nibbleArray) Set(i int, nu int) {}
