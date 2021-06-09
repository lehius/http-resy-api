package config

import "gopkg.in/yaml.v2"

// Config ...
type Config struct {
	Network        *Network `yaml:"network"`
	Logging        *Logging `yaml:"logging"`
	configFilePath string
}

// Network ...
type Network struct {
	Bind string `yaml:"bind"`
}

// Logging ...
type Logging struct {
	Level string `yaml:"level"`
}

// NewConfig ...
func Load() (*Config, error) {
	result := Default()
	cfg, err := result.ReadFile()
	if err != nil {
		return nil, err
	}
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, result)
	return result, err
}
