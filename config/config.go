package config

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

var Config *yaml.File

func Initial(yamlFile string){
	var err error
	Config, err = yaml.ReadFile(yamlFile)
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	value, err := Config.Get(key)
	if err != nil {
		panic(err)
	}
	return value
}

func GetInt(key string) int {
	value, err := Config.GetInt(key)
	if err != nil {
		panic(err)
	}
	return int(value)
}

func GetBool(key string) bool {
	value, err := Config.GetBool(key)
	if err != nil {
		panic(err)
	}
	return value
}