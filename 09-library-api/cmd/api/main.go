package main

import (
	"go-lab/09-library-api/internal/handlers"
	"go-lab/09-library-api/internal/server"
	"log"
)

func main() {
	homeHandler := handlers.NewHomeHandler()
	bookHandler := handlers.NewBookHandler()

	srv := server.New(homeHandler, bookHandler)
	
	log.Println("starting library api...")

	if err := srv.Start() ; err != nil {
		log.Fatal(err)
	}
}