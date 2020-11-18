package config

import (
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetDefault(HTTPServicePort, "8080")
	viper.SetDefault(GRPCServicePort, "5000")
	viper.SetConfigFile("local-config.yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv() // read from env variable
}
