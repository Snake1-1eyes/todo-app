package main

import (
	"log"

	"github.com/Snake1-1eyes/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
