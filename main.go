package main

import (
	"fmt"
	"net/http"

	"github.com/gunturaf/omdb-server/config"
	"github.com/gunturaf/omdb-server/controllers/grpcservice"
	"github.com/gunturaf/omdb-server/controllers/httpapi"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/spf13/viper"
)

func main() {
	httpClient := http.DefaultClient

	config.ReadConfig()

	omdbService := omdbservice.NewOMDBService(httpClient, viper.GetString(config.OMDBApiBaseURL), viper.GetString(config.OMDBApiKey))

	go httpapi.RunServer(viper.GetString(config.HTTPServicePort), omdbService)
	fmt.Println("http api running at :" + viper.GetString(config.HTTPServicePort))

	go grpcservice.RunGRPCServer(viper.GetString(config.GRPCServicePort), omdbService)
	fmt.Println("grpc server running at :" + viper.GetString(config.GRPCServicePort))

	select {}
}
