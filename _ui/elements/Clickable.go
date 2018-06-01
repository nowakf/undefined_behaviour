package elements

type Clickable interface {
	OnMouse(clicked bool)
	Flush()
	Do()
}
