package config

import (
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

// Config 申明全局配置类型
type Config struct {
	DBURI    string `yaml:"DBURI"`
	MongoURI string `yaml:"MONGO_URI"`
}

// Cfg 全局配置
var Cfg *Config

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

// InitConfig 初始化全局配置
func InitConfig() error {
	result, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	confContent := replaceEnvInConfig(result)

	if err := yaml.Unmarshal(confContent, &Cfg); err != nil {
		return err
	}

	return nil
}
