package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/hashicorp/vault-client-go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// var (
// 	GitUrl     string
// 	GitCommit  string
// 	JenkinsJob string
// )

// func LoadAbout() *About {
// 	return &About{
// 		Git: GitConfig{
// 			Repository: GitUrl,
// 			CommitId:   GitCommit,
// 		},
// 		Jenkins: JenkinsConfig{
// 			Job: JenkinsJob,
// 		},
// 	}
// }

type Config struct {
	App      APPConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Server   ServerConfig   `mapstructure:"server"`
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

type ServerConfig struct {
	ExpireToken string `mapstructure:"expire_token"`
	HostKey     string `mapstructure:"host_key"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Allow environment variables to override config values
	// v.AutomaticEnv()
	// v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

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
	// if err := v.ReadInConfig(); err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		return nil, errors.New("config file not found")
	// 	}
	// 	return nil, err
	// }
	// log.Println("Loaded config from config.yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Allow environment variables to override config values
			v.AutomaticEnv()
			v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

			// auto bind environment variables with reflection
			if err := autoBindEnv(v, Config{}); err != nil {
				log.Printf("Warning during auto binding: %v", err)
			}
		} else {
			return nil, errors.New("config file not found")
		}
	} else {
		log.Println("Loaded config from config.yaml")
	}

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

// func loadConfigFromVault(config *Config) error {
// 	ctx := context.Background()

// 	// Create Vault client
// 	client, err := vault.New(
// 		vault.WithEnvironment(),
// 		vault.WithTLS(vault.TLSConfiguration{
// 			InsecureSkipVerify: true,
// 		}),
// 		vault.WithRequestTimeout(30*time.Second),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to create Vault client: %w", err)
// 	}

// 	// Set the token
// 	if err := client.SetToken(os.Getenv("VAULT_TOKEN")); err != nil {
// 		return fmt.Errorf("failed to set Vault token: %w", err)
// 	}

// 	// Load Database configuration
// 	if err := loadSectionFromVault(ctx, client, "Dev", "apps/production", &config.Database); err != nil {
// 		log.Printf("Warning: failed to load Database config from Vault: %v", err)
// 	}
// 	config.Database.Host = "localhost"

// 	// Load Database configuration
// 	if err := loadSectionFromVault(ctx, client, "Ops", "apps/production", &config.Database); err != nil {
// 		log.Printf("Warning: failed to load Database config from Vault: %v", err)
// 	}

// 	return nil
// }

// func loadSectionFromVault(ctx context.Context, client *vault.Client, mountPath string, secretPath string, target interface{}) error {
// 	// Read secret from KV v2
// 	secret, err := client.Secrets.KvV2Read(ctx, secretPath, vault.WithMountPath(mountPath))
// 	if err != nil {
// 		return fmt.Errorf("failed to read secret from %s: %w", secretPath, err)
// 	}

// 	if secret == nil || secret.Data.Data == nil {
// 		return fmt.Errorf("no data found at %s", secretPath)
// 	}

// 	// Use mapstructure to decode the data into the target struct
// 	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
// 		TagName:          "mapstructure",
// 		Result:           target,
// 		WeaklyTypedInput: true,  // Allow type conversion
// 		ErrorUnused:      false, // Don't error on unused fields
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to create decoder: %w", err)
// 	}

// 	if err := decoder.Decode(secret.Data.Data); err != nil {
// 		return fmt.Errorf("failed to decode data: %w", err)
// 	}

// 	return nil
// }

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

func autoBindEnv(v *viper.Viper, cfg interface{}, prefix ...string) error {
	cfgType := reflect.TypeOf(cfg)
	cfgValue := reflect.ValueOf(cfg)

	// Handle pointers
	if cfgType.Kind() == reflect.Ptr {
		cfgType = cfgType.Elem()
		cfgValue = cfgValue.Elem()
	}

	// Only process structs
	if cfgType.Kind() != reflect.Struct {
		return nil
	}

	// Base path for nested keys
	basePath := ""
	if len(prefix) > 0 {
		basePath = prefix[0]
	}

	// Process each field in the struct
	for i := 0; i < cfgType.NumField(); i++ {
		field := cfgType.Field(i)
		fieldValue := cfgValue.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Get the key name from mapstructure tag or field name
		key := field.Tag.Get("mapstructure")
		if key == "" {
			key = strings.ToLower(field.Name)
		}

		// Build the full config path
		configPath := key
		if basePath != "" {
			configPath = basePath + "." + key
		}

		// Handle nested structs recursively
		if field.Type.Kind() == reflect.Struct {
			if err := autoBindEnv(v, fieldValue.Interface(), configPath); err != nil {
				return err
			}
			continue
		}

		// Convert config path to environment variable name
		envVar := strings.ToUpper(strings.ReplaceAll(configPath, ".", "_"))

		// Bind environment variable to config path
		if err := v.BindEnv(configPath, envVar); err != nil {
			return fmt.Errorf("failed to bind env var %s: %w", envVar, err)
		}
	}

	return nil
}
