package main

import (
	"net/http"

	"github.com/gunturaf/omdb-server/config"
	"github.com/gunturaf/omdb-server/controllers/httpapi"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/spf13/viper"
)

func main() {
	httpClient := http.DefaultClient

	config.ReadConfig()

	omdbService := omdbservice.NewOMDBService(httpClient, viper.GetString(config.OMDBApiBaseURL), viper.GetString(config.OMDBApiKey))

	httpapi.RunServer(viper.GetString(config.HTTPServicePort), omdbService)
}
