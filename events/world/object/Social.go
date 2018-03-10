package object

type Social struct {
	relationships []relationship
}

type relationship struct {
	entity       int
	state        int
	descriptions []string
	deltas       []int
}
