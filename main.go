package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gunturaf/omdb-server/config"
	"github.com/gunturaf/omdb-server/controllers/grpcservice"
	"github.com/gunturaf/omdb-server/controllers/httpapi"
	"github.com/gunturaf/omdb-server/infrastructure/repository/mysqldb"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
	"github.com/spf13/viper"
)

func init() {
	config.ReadConfig()
}

func connectMysqlDB() *sql.DB {
	mysqlDSL := mysqldb.MysqlDBDSL{
		Username: viper.GetString(config.MysqlUsername),
		Password: viper.GetString(config.MysqlPassword),
		Host:     viper.GetString(config.MysqlHost),
		Port:     viper.GetString(config.MysqlPort),
		DBName:   viper.GetString(config.MysqlDBName),
	}

	db, err := sql.Open("mysql", mysqlDSL.GetDSN())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func main() {
	httpClient := http.DefaultClient

	mysqlRepo := mysqldb.NewMysqlDB(connectMysqlDB())
	omdbService := omdbservice.NewOMDBService(httpClient, viper.GetString(config.OMDBApiBaseURL), viper.GetString(config.OMDBApiKey))
	searchUseCase := usecase.NewSearchUseCase(omdbService, mysqlRepo)

	go httpapi.RunServer(viper.GetString(config.HTTPServicePort), omdbService, searchUseCase)
	fmt.Println("http api running at :" + viper.GetString(config.HTTPServicePort))

	go grpcservice.RunGRPCServer(viper.GetString(config.GRPCServicePort), omdbService, searchUseCase)
	fmt.Println("grpc server running at :" + viper.GetString(config.GRPCServicePort))

	select {}
}
