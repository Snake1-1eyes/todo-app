package main

import (
	"log"

	"github.com/Snake1-1eyes/todo-app"
	"github.com/Snake1-1eyes/todo-app/pkg/handler"
	"github.com/Snake1-1eyes/todo-app/pkg/repository"
	"github.com/Snake1-1eyes/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
