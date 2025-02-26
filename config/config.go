package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/vault-client-go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	GitUrl     string
	GitCommit  string
	JenkinsJob string
)

func LoadAbout() *About {
	return &About{
		Git: GitConfig{
			Repository: GitUrl,
			CommitId:   GitCommit,
		},
		Jenkins: JenkinsConfig{
			Job: JenkinsJob,
		},
	}
}

type Config struct {
	App      APPConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type APPConfig struct {
	Port string `mapstructure:"url"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Allow environment variables to override config values
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if checkVaultEnv() {
		if err := loadConfigFromVault(&cfg); err != nil {
			return nil, err
		}
		return &cfg, nil
	}

	// Set the paths and config name/type
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	log.Println("Loaded config from config.yaml")

	// unmarshal the configuratioon into the Config struct
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func loadConfigFromVault(config *Config) error {
	ctx := context.Background()
	client, err := vault.New(
		vault.WithEnvironment(),
		vault.WithTLS(vault.TLSConfiguration{
			InsecureSkipVerify: true,
		}))
	if err != nil {
		log.Fatal(err)
	}

	secret, err := client.Secrets.KvV2Read(ctx, "/foo/bar", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}

	if secret == nil || secret.Data.Data == nil {
		return errors.New("no data found at the specified Vault path")
	}

	if err := mapstructure.Decode(secret.Data.Data, &config); err != nil {
		return fmt.Errorf("error mapping data to struct: %s", err)
	}
	log.Println("Loaded secrets from Vault and assigned to config")

	return nil
}

func checkVaultEnv() bool {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")

	if vaultAddr == "" && vaultToken == "" {
		return false
	}

	// If only one of them is set, return an error
	if (vaultAddr != "" && vaultToken == "") || (vaultAddr == "" && vaultToken != "") {
		missing := "VAULT_TOKEN"
		if vaultAddr == "" {
			missing = "VAULT_ADDR"
		}
		log.Fatalf("incomplete Vault configuration: %s is missing", missing)
	}
	return true
}
