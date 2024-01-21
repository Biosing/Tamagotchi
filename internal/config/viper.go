package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile("config/.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
}
