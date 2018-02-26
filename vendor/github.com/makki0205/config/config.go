package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Yaml = getYaml()
var goEnv = getGoEnv()

func Env(key string) string {
	return Yaml[goEnv].(map[interface{}]interface{})[key].(string)
}

func getYaml() map[interface{}]interface{} {
	var buf, err = ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func getGoEnv() string {
	_goEnv := os.Getenv("GO_ENV")
	if _goEnv == "" {
		return "develop"
	}
	return _goEnv
}
