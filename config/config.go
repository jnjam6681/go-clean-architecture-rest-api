package config

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      APPConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type APPConfig struct {
	Port string `mapstructure:"port"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

var c Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	// var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
