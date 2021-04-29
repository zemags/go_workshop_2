package main

import (
	"log"

	"github.com/spf13/viper"
	workshop_2 "github.com/zemags/go_workshop_2"
	"github.com/zemags/go_workshop_2/pkg/handler"
	"github.com/zemags/go_workshop_2/pkg/repository"
	"github.com/zemags/go_workshop_2/pkg/service"
)

func main() {
	if err := ininConfig(); err != nil {
		log.Fatalf("error loading configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(workshop_2.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func ininConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
