package main

import (
	todo "Todo"
	"Todo/pkg/handler"
	"Todo/pkg/repository"
	"Todo/pkg/service"
	// "fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("err initializing configs:%s", err.Error())
		// fmt.Println(err)

	}
	repos := repository.NewRepository()
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yml")
	return viper.ReadInConfig()
}
