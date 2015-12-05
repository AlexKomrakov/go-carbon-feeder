package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Interval int
	Command
}

func (c Config) FromYml(file string) (Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(content, &c)

	return c, err
}