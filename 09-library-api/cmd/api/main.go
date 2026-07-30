package main

import (
	"go-lab/09-library-api/internal/config"
	"go-lab/09-library-api/internal/database"
	"go-lab/09-library-api/internal/handlers"
	"go-lab/09-library-api/internal/repositories"
	"go-lab/09-library-api/internal/server"
	"go-lab/09-library-api/internal/services"
	"log"
)

func main() {
	cfg := config.Load()
	database,err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	bookRepo := repositories.NewPostgressBookRepository(database)
	bookService := services.Newbookservice(bookRepo)

	homeHandler := handlers.NewHomeHandler()
	bookHandler := handlers.NewBookHandler(bookService)

	srv := server.New(homeHandler, bookHandler)
	
	log.Println("starting library api...")

	if err := srv.Start() ; err != nil {
		log.Fatal(err)
	}
}