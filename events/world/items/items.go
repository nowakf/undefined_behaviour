package items

import . "github.com/nowakf/undefined_behaviour/events/world/object"

type Item Key

const (
	Gun         Item = Item(BOOL0)
	Car              = Item(BOOL1)
	Uniform          = Item(NIBBLE16)
	PlaceOfWork      = Item(NIBBLE20)
	Money            = Item(UINT8_48)
)

type Items Object

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
