package config

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      APPConfig
	Postgres PostgresConfig
}

type APPConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  bool
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set the paths and config name/type
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Allow environment variables to override config values
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	// unmarshal the configuratioon into the Config struct
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}
