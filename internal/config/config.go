package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"rest-api-crud/pkg/logging"
	"sync"
	"time"
)

type (
	LoggingConfig struct {
		LogLevel  logging.LogLevel  `yaml:"log-level" env:"LOG_LEVEL" env-default:"info"`
		LogFormat logging.LogFormat `yaml:"log-format" env:"LOG_FORMAT" env-default:"text"`
	}
	ListenConfig struct {
		Type   string `yaml:"type" env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `yaml:"bind_ip" env:"LISTEN_BIND_IP" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env:"LISTEN_PORT" env-default:"8080"`
	}
	MongoDBConfig struct {
		Host       string `yaml:"host" env:"MONGO_HOST" env-required:"true"`
		Port       string `yaml:"port" env:"MONGO_PORT" env-required:"true"`
		Database   string `yaml:"database" env:"MONGO_DATABASE" env-required:"true"`
		AuthDB     string `yaml:"auth_db" env:"MONGO_AUTH_DB"`
		Username   string `yaml:"username" env:"MONGO_USERNAME"`
		Password   string `yaml:"password" env:"MONGO_PASSWORD"`
		Collection string `yaml:"collection" env:"MONGO_COLLECTION"`
	}
	PostgresConfig struct {
		Host              string        `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
		Port              string        `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
		Database          string        `yaml:"database" env:"POSTGRES_DATABASE" env-required:"true"`
		User              string        `yaml:"username" env:"POSTGRES_USER" env-required:"true"`
		Password          string        `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
		ConnRetryAttempts int           `yaml:"conn_retry_attempts" env:"POSTGRES_CONN_RETRY_ATTEMPTS" env-default:"5"`
		ConnRetryDelay    time.Duration `yaml:"conn_retry_delay" env:"POSTGRES_CONN_RETRY_DELAY" env-default:"5s"`
	}
	Config struct {
		IsDebug  bool           `yaml:"is_debug" env:"IS_DEBUG" env-required:"true"`
		Logging  LoggingConfig  `yaml:"logging"`
		Listen   ListenConfig   `yaml:"listen"`
		MongoDB  MongoDBConfig  `yaml:"mongodb"`
		Postgres PostgresConfig `yaml:"postgres"`
	}
)

var instance *Config
var once sync.Once

func GetConfig() *Config {
	logger := logging.GetLogger()

	logger.Info("Loading configuration")

	once.Do(func() {
		instance = &Config{}
		help, _ := cleanenv.GetDescription(instance, nil)
		if err := cleanenv.ReadConfig("configs/config.yml", instance); err != nil {
			logger.Error("Failed to load configuration",
				"error", err,
				"help", help)
			panic("Configuration loading failed")
		}
		logger.Info("Configuration loaded successfully")
	})
	return instance
}
