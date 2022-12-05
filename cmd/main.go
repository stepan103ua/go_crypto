package main

import (
	crypto_app "crypto_app/pkg"
	handler "crypto_app/pkg/handler"
	"crypto_app/pkg/repository"
	"crypto_app/pkg/service"
	"log"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	configError := initConfigs()
	if configError != nil {
		log.Fatalf("Error ocurred on init config: %s", configError.Error())
	}

	db, dbError := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if dbError != nil {
		log.Fatalf("Error ocurred on connecting to database: %s", dbError.Error())
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := new(crypto_app.Server)
	err := server.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("Error ocurred on run: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
