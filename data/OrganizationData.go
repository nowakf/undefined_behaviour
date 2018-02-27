package data

type OrganizationDataData struct {
	Name        string         `yaml:"Name"`
	Tags        string         `yaml:"tags"`
	Description string         `yaml:"description"`
	Goals       []GoalData     `yaml:"goals"`
	Effects     map[string]int `yaml:"effects"`
}

type GoalData struct {
	Name        string `yaml:"Name"`
	Tags        string `yaml:"tags"`
	Description string `yaml:"description"`
	Alignment   int    `yaml:"alignment"`
}
