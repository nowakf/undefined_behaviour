package events

type Event struct {
	Event_url    string   `yaml:"event_url"`
	Email        Email    `yaml:"email"`
	Article      Article  `yaml:"article"`
	Consequences []string `yaml:"consequences"`
}
