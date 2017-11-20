package config

import (
	"io/ioutil"
	. "crypto-arbitrage/services/arg-parser"
	. "crypto-arbitrage/helpers"
	"fmt"
	"gopkg.in/yaml.v2"
)

var Config configModel

type configModel struct {
	App struct {
		Env   string
		Debug bool
	}
	Storage struct {
		Postgres struct {
			Host string
			Port string
			Name string
			User string
			Pass string
		}
	}
}

func (t *configModel) readFile(env string) []byte {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("./config/%s.yaml", env))
	Error.Check(err)
	return yamlFile
}

func (t *configModel) Load() {
	file := t.readFile(ArgumentParser.Config.Value)
	err := yaml.Unmarshal(file, t)
	Error.Check(err)
}



func init() {
	if Config == (configModel{}) {
		Config = *new(configModel)
		Config.Load()
	}
}
