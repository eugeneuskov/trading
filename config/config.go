package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Exchange struct {
	Name      string `yaml:"exchange_name"`
	ApiKey    string `yaml:"api_key"`
	ApiSecret string `yaml:"api_secret"`
}

type Config struct {
	Exchanges []Exchange
}

func (c *Config) Init() (*Config, error) {
	file, err := os.Open("config/config.yml")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, err
}
