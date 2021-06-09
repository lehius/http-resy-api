package config

import (
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

func (c *Config) ReadFile() (*Config, error) {
	cfg := &Config{}
	err := read(c.configFilePath, cfg)
	return cfg, err
}

func read(filename string, r interface{}) (err error) {
	file, err := ioutil.ReadFile(path.Clean(filename))
	if err == nil {
		err = yaml.Unmarshal(file, r)
	}
	return
}
