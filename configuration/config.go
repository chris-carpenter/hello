package configuration

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	helloMessage string
}

func Config() {
	//environment
	viper.SetEnvPrefix("hello")
	viper.BindEnv("battle_client_id")
	viper.BindEnv("battle_client_secret")
	//application.yml
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

