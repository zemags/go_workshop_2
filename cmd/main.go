package main

import (
	"log"

	workshop_2 "github.com/zemags/go_workshop_2"
	"github.com/zemags/go_workshop_2/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(workshop_2.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
