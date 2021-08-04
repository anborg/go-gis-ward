package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	App AppConfig `yaml:"app"`
}

// AppConfig App related config - eg logging
type AppConfig struct {
	Port        string `yaml:"port"`
	WardGeoJson string `yaml:"ward_geojson"`
}

func (cfg *Config) New(path string) (err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return err
	}
	return
}
