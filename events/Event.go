package events

type Event struct {
	Event_url    string   `yaml:"event_url"`
	Etype        etype    `yaml:"etype"`
	Title        string   `yaml:"title"`
	Content      string   `yaml:"content"`
	Consequences []string `yaml:"consequences"`
}
