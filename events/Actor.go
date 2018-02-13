package events

type Actor struct {
	Name          string
	Org           Organization
	MailBox       Event
	Abilities     map[bool]Action
	Relationships map[int]Actor
}
