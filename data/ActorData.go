package data

type ActorData struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	Tags         []string `yaml:"tags"`
	Alignment    string   `yaml:"alignment"`
	MessageRoot  string   `yaml:"message_root"`
	Organization string   `yaml:"organization"`
}
