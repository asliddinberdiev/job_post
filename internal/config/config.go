package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App     App
	MongoDB MongoDB
}

type App struct {
	Environment string `envconfig:"APP_ENVIRONMENT" default:"dev" required:"true"`
	Name        string `envconfig:"APP_NAME" default:"job_post" required:"true"`
	Host        string `envconfig:"APP_HOST" default:"localhost" required:"true"`
	Port        int    `envconfig:"APP_PORT" default:"8000" required:"true"`
}

type MongoDB struct {
	Host     string `envconfig:"MONGODB_HOST" default:"localhost" required:"true"`
	Port     int    `envconfig:"MONGODB_PORT" default:"27017" required:"true"`
	Username string `envconfig:"MONGODB_USERNAME" default:"username" required:"true"`
	Password string `envconfig:"MONGODB_PASSWORD" default:"password" required:"true"`
	Database string `envconfig:"MONGODB_DATABASE" default:"job_post" required:"true"`
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	var cfg Config

	if err := envconfig.Process("APP", &cfg.App); err != nil {
		return nil, err
	}
	if err := envconfig.Process("MONGODB", &cfg.MongoDB); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) GetAppAddress() string {
	return fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)
}

func (c *Config) GetMongoURL() string {
	return fmt.Sprintf("mongodb://%s:%d", c.MongoDB.Host, c.MongoDB.Port)
}
