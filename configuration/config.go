package configuration

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	helloMessage string
}

func Config() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Unable to find config file, %s", err)
		} else {
			log.Fatalf("Error getting config, %s", err)
		}
	}
}

