package config

import "github.com/spf13/viper"

func ReadConfig() {
	viper.SetDefault("HTTP_SERVICE_PORT", "8080")
	viper.SetConfigFile("local-config.yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv() // read from env variable
}
