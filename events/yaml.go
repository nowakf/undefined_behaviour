package events

import (
	"gopkg.in/yaml.v2"
	"log"

	"os"

	"io/ioutil" //this should be moved
)

func load() []Event {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/assets/text/test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	e, err := fromYAML(yamlFile)
	if err != nil {
		log.Fatalf("Unmarshal %v", err)
	}
	return e
}

func fromYAML(yamlFile []byte) ([]Event, error) {

	var e []Event

	err := yaml.UnmarshalStrict(yamlFile, &e)

	return e, err

}
