package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl  string `mapstructure:"DB_URI"`
	DBName string `mapstructure:"DB_NAME"`
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
}

var envs = []string{
	"DB_URI", "DB_NAME", "HOST", "PORT",
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
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
