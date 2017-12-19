package configuration

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func (config *ConfigFile) Load(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Configfile load read error: %v", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Configfile load unmarshal error: %v", err)
		return err
	}
	return nil
}
