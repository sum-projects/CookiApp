package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Domain        string `mapstructure:"MAIL_DOMAIN"`
	Host          string `mapstructure:"MAIL_HOST"`
	Port          int    `mapstructure:"MAIL_PORT"`
	Username      string `mapstructure:"MAIL_USERNAME"`
	Password      string `mapstructure:"MAIL_PASSWORD"`
	Encryption    string `mapstructure:"MAIL_ENCRYPTION"`
	FromName      string `mapstructure:"MAIL_FROM_NAME"`
	FromAddress   string `mapstructure:"MAIL_FROM_ADDRESS"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
