package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port     int    `json:"port"`
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfig(fname string) (*Config, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}
