package object

var descriptions = [][]description{}

type description struct {
}

func (d *description) Title() string {
	return ""
}

func (d *description) Blurb() string {
	return ""
}
