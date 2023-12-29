package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenPort string `envconfig:"HTTP_PORT"`
	MongoURL   string `envconfig:"MONGO_URI_ORDERROUTE"`
	Database   string `envconfig:"DATABASE"`
	JwtSecret  string `envconfig:"JWT_SECRET"`
}

func GetConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
