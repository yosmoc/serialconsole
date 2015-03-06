package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Serial SerialConfig `yaml:"serial"`
}

type SerialConfig struct {
	Port string `yaml:"port"`
	Baud int    `yaml:"baudrate"`
}

func ParseConf(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	if c.Serial.Port == "" {
		return nil, errors.New("Port config is needed")
	}

	if c.Serial.Baud == 0 {
		return nil, errors.New("Baudrate config is needed")
	}

	return c, nil
}
