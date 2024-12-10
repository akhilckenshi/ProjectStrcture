package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl     string `mapstructure:"DB_URI"`
	DBName    string `mapstructure:"DB_NAME"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	Port      string `mapstructure:"PORT"`
}

var envs = []string{
	"DB_URI", "DB_NAME", "JWT_SECRET", "PORT",
}

func LoadConfig() (Config, error) {
	var config Config

	// Load the .env file if it exists
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read .env file: %v\n", err)
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
