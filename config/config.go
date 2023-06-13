package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	appModeProd = "prod"
)

type Exchange struct {
	Id        string   `yaml:"exchange_id"`
	ApiKey    string   `yaml:"api_key"`
	ApiSecret string   `yaml:"api_secret"`
	Url       []string `yaml:"exchange_url"`
}

type Config struct {
	Exchanges []Exchange
}

func (c *Config) Init() (*Config, error) {
	applicationMode, _ := envValue("APP_ENV")

	file, err := os.Open(configFilePath(applicationMode))
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

func configFilePath(applicationMode string) string {
	if applicationMode == appModeProd {
		return "config/config_prod.yml"
	}

	return "config/config_test.yml"
}

func envValue(envKey string) (string, error) {
	value, found := os.LookupEnv(envKey)
	if found && value != "" {
		return value, nil
	}

	reader, err := os.Open(".env")
	if err != nil {
		return "", err
	}
	defer func(reader *os.File) {
		_ = reader.Close()
	}(reader)

	env, err := godotenv.Parse(reader)
	value, found = env[envKey]
	if !found {
		return "", fmt.Errorf("%s not found", envKey)
	}

	if value == "" {
		return "", fmt.Errorf("%s is empty", envKey)
	}

	return value, nil
}
