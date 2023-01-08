package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadYAMLFile(filename string, out interface{}) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlFile, out)
}
