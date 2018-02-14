package events

type Actor struct {
	Name          string
	Org           Organization
	MailBox       Event
	Abilities     map[bool]Action
	Relationships map[int]Actor
}

//GenerateActor is for when you want to randomly generate a person, for
//whom no content exists.
func GenerateActor(w *world) Actor {
	return *new(Actor)
}
