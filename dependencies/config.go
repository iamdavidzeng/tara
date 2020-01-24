package dependencies

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
)

type Config struct {
	DBURI string `yaml:"DBURI"`
	MONGO_URI string `yaml:"MONGO_URI"`
}

var config *Config

func replaceEnvInConfig(body []byte) []byte {
	search := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	replacedBody := search.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := search.ReplaceAllString(string(b), `$1`)
		group2 := search.ReplaceAllString(string(b), `$2`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte(group2)
	})

	return replacedBody
}

func InitConfig() error {
	result, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	confContent := replaceEnvInConfig(result)

	if err := yaml.Unmarshal(confContent, &config); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}
