package dependencies

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DBURI string `yaml:"DBURI"`
}

var config *Config

func InitConfig() error {
	result, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	_config := yaml.Unmarshal(result, &config)

	return _config
}

func GetConfig() *Config {
	return config
}