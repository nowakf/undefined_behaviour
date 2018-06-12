package data

import (
	"gopkg.in/yaml.v2"

	"os"
)

type yamlLoader struct {
	path string
}

func NewYamlLoader(path string) *yamlLoader {
	y := new(yamlLoader)
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	y.path = pwd + path

	return y
}

func (y *yamlLoader) StoriesFromYaml(yamlFile []byte) ([]StoryData, error) {

	var e []StoryData

	err := yaml.UnmarshalStrict(yamlFile, &e)

	return e, err

}

func (y *yamlLoader) AbilitiesFromYaml(yamlFile []byte) ([]AbilityData, error) {

	var e []AbilityData

	err := yaml.UnmarshalStrict(yamlFile, &e)

	return e, err

}

func (y *yamlLoader) MessagesFromYaml(yamlFile []byte) ([]MessageData, error) {

	var e []MessageData

	err := yaml.UnmarshalStrict(yamlFile, &e)

	return e, err

}
