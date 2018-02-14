package events

type Organization struct {
	EmailRoot string
	members   []*Actor
}
