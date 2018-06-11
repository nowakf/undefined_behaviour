package world

import o "ub/events/world/object"

type Person struct {
	Data   *o.Object
	id     o.Identity
	body   o.Body
	mind   o.Mind
	social o.Social
}

func NewPerson() *Person {
	p := new(Person)
	return p
}
func (p *Person) Queue(url string) {}
func (p *Person) Body() *o.Body {
	return &p.body
}

func (p *Person) whatever() {
	print(p.Body().Strength())
}
func (p *Person) Mind() *o.Mind {
	return &p.mind
}

func (p *Person) Social() *o.Social {
	return &p.social
}

func (p *Person) ID() *o.Identity {
	return &p.id
}
