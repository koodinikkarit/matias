package configuration

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func (config *ConfigFile) Save(path string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Configfile save marshal error: %v", err)
		return err
	}
	err = ioutil.WriteFile(path, data, os.ModeCharDevice)
	if err != nil {
		log.Fatalf("Configfile save write error: %v", err)
		return err
	}
	return nil
}
