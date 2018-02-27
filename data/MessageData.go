package data

type MessageData struct {
	Title   string         `yaml:"title"`
	Tags    string         `yaml:"tags"`
	Content string         `yaml:"content"`
	Options []string       `yaml:"options"`
	Effects map[string]int `yaml:"effects"`
}
