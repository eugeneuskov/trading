package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	appModeProd = "prod"

	postgresHostEnvKey     = "POSTGRES_HOST"
	postgresPortEnvKey     = "POSTGRES_PORT"
	postgresUserEnvKey     = "POSTGRES_USER"
	postgresPasswordEnvKey = "POSTGRES_PASSWORD"
	postgresDatabaseEnvKey = "POSTGRES_DATABASE"
)

type Exchange struct {
	Id        string   `yaml:"exchange_id"`
	ApiKey    string   `yaml:"api_key"`
	ApiSecret string   `yaml:"api_secret"`
	Url       []string `yaml:"exchange_url"`
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type Config struct {
	Exchanges []Exchange
	Db
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

	dbHost, _ := envValue(postgresHostEnvKey)
	dbPort, _ := envValue(postgresPortEnvKey)
	dbUser, _ := envValue(postgresUserEnvKey)
	dbPassword, _ := envValue(postgresPasswordEnvKey)
	dbName, _ := envValue(postgresDatabaseEnvKey)

	config.Db = Db{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DbName:   dbName,
		SslMode:  "disable",
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
