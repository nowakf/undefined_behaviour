package data

type AbilityData struct {
	Title        string         `yaml:"title"`
	Description  string         `yaml:"description"`
	Tags         []string       `yaml:"tags"`
	Effects      map[string]int `yaml:"effects"`
	Requirements map[string]int `yaml:"requirements"`
	Progression  string         `yaml:"progression"`
}
