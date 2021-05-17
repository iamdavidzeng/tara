package config

import (
	"os"
	"regexp"

	"github.com/spf13/viper"
)

type Config struct {
	DB *DB
}

type DB struct {
	URL string `mapstructure:"DB_URL"`
}

var Cfg *Config

func (c *Config) Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	for _, key := range viper.AllKeys() {
		viper.Set(key, parse(viper.GetString(key)))
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	return nil
}

func parse(s string) string {
	compiler := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	value := compiler.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		match := compiler.FindStringSubmatch(string(b))
		envValue, defaultValue := os.Getenv(match[1]), match[2]
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte(defaultValue)

	})
	return string(value)
}
