package items

import . "github.com/nowakf/object"

var List = []Item{
	Gun,
	Car,
	Uniform,
	PlaceOfWork,
	Money}

const (
	IMAX = Items(MAX)
)

type Item Key

const (
	Gun         Item = Item(BOOL0)
	Car              = Item(BOOL1)
	Uniform          = Item(NIBBLE16)
	PlaceOfWork      = Item(NIBBLE20)
	Money            = Item(UINT8_48)
)

type Items Object

func (i Items) Superset(this Items) bool {
	return Object(i).Superset(Object(this))
}
func (i Items) Count() int {
	return len(Object(i).Fields())
}

func (i Items) Fields() (vals []Item) {
	fields := Object(i).Fields()
	vals = make([]Item, len(fields))
	for i, field := range fields {
		vals[i] = Item(field)
	}
	return
}

//returns if you have an item, and how many
func (i Items) Has(a Item) int {
	return Object(i).Get(Key(a))
}

func (i *Items) Set(a Item, newVal int) *Items {
	(*Object)(i).Set(Key(a), newVal)
	return i
}

func (i Item) Describe() (description string) {
	switch i {
	case Gun:
		return "a pistol"
	case Car:
		return "a car"
	case Money:
		return "money"
	case Uniform:
		return "a uniform"
	case PlaceOfWork:
		return "is employed here"
	default:
		return ""
	}
}
