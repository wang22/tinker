package config

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

var config *yaml.File

func Initial(yamlFile string) {
	var err error
	config, err = yaml.ReadFile(yamlFile)
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	value, err := config.Get(key)
	if err != nil {
		panic(err)
	}
	return value
}

func GetInt(key string) int {
	value, err := config.GetInt(key)
	if err != nil {
		panic(err)
	}
	return int(value)
}

func GetBool(key string) bool {
	value, err := config.GetBool(key)
	if err != nil {
		panic(err)
	}
	return value
}
