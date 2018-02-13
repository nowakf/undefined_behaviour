package events

type Email struct {
	Subject string `yaml:"subject"`
	Sender  string `yaml:"sender"`
	Content string `yaml:"content"`
}
