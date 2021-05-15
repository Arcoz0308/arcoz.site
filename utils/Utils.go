package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Token string `yaml:"token"`
}

var Config config

func LoadConfig() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal()
	}
}
