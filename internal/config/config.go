package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"rest-api-crud/pkg/logging"
	"sync"
	"time"
)

type (
	ListenConfig struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	}
	MongoDBConfig struct {
		Host       string `yaml:"host" env-required:"true"`
		Port       string `yaml:"port" env-required:"true"`
		Database   string `yaml:"database" env-required:"true"`
		AuthDB     string `yaml:"auth_db"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	}
	PostgresConfig struct {
		Host              string        `yaml:"host" env-default:"localhost"`
		Port              string        `yaml:"port" env-default:"5432"`
		Database          string        `yaml:"database" env-required:"true"`
		User              string        `yaml:"username" env-required:"true"`
		Password          string        `yaml:"password" env-required:"true"`
		ConnRetryAttempts int           `yaml:"conn_retry_attempts" env-default:"5"`
		ConnRetryDelay    time.Duration `yaml:"conn_retry_delay" env-default:"5s"`
	}
	Config struct {
		IsDebug  bool           `yaml:"is_debug" env-required:"true"`
		Listen   ListenConfig   `yaml:"listen"`
		MongoDB  MongoDBConfig  `yaml:"mongodb"`
		Postgres PostgresConfig `yaml:"postgres"`
	}
)

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Reading application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("configs/config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Warn(help)
			logger.Fatal(err)
		}
		logger.Info("Configuration parsed successfully")
	})
	return instance
}
