package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	App     App     `mapstructure:"app"`
	MongoDB MongoDB `mapstructure:"mongodb"`
}

type App struct {
	Environment string `envconfig:"APP_ENVIRONMENT" default:"dev" required:"true" mapstructure:"environment"`
	Name        string `envconfig:"APP_NAME" default:"job_post" required:"true" mapstructure:"name"`
	Host        string `envconfig:"APP_HOST" default:"localhost" required:"true" mapstructure:"host"`
	Port        int    `envconfig:"APP_PORT" default:"8000" required:"true" mapstructure:"port"`
}

type MongoDB struct {
	Host     string `envconfig:"MONGODB_HOST" default:"localhost" required:"true"`
	Port     int    `envconfig:"MONGODB_PORT" default:"27017" required:"true"`
	Username string `envconfig:"MONGODB_USERNAME" default:"username" required:"true"`
	Password string `envconfig:"MONGODB_PASSWORD" default:"password" required:"true"`
	Database string `envconfig:"MONGODB_DATABASE" default:"job_post" required:"true"`
}

func Init(confDir string) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	if confDir == "" {
		v.AddConfigPath("./config")
	}
	v.AddConfigPath(confDir)

	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	v.SetConfigName(env)

	envFile := fmt.Sprintf("%s/%s.env", confDir, env)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: No %s.env file found in %s: %v", env, confDir, err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	if v.ConfigFileUsed() != "" {
		if err := v.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func (c *Config) GetAppAddress() string {
	return fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)
}

func (c *Config) GetMongoURL() string {
	return fmt.Sprintf("mongodb://%s:%d", c.MongoDB.Host, c.MongoDB.Port)
}
