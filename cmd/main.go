package main

import (
	"log"

	"github.com/Dorkhan1/todo-app"
	"github.com/Dorkhan1/todo-app/pkg/handler"
	"github.com/Dorkhan1/todo-app/pkg/repository"
	"github.com/Dorkhan1/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error inintializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "ToDo-db",
		SSlMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize bd:%s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
