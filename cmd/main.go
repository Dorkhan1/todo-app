package main

import (
	"log"

	"github.com/Dorkhan1/todo-app"
	"github.com/Dorkhan1/todo-app/pkg/handler"
	"github.com/Dorkhan1/todo-app/pkg/repository"
	"github.com/Dorkhan1/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}


}
