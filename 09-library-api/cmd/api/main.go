package main

import (
	"go-lab/09-library-api/internal/config"
	"go-lab/09-library-api/internal/database"
	"go-lab/09-library-api/internal/handlers"
	"go-lab/09-library-api/internal/repositories"
	"go-lab/09-library-api/internal/server"
	"go-lab/09-library-api/internal/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	database,err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	bookRepo := repositories.NewPostgressBookRepository(database)
	userRepo := repositories.NewPostgressUserRepository(database)
	borrowRepo := repositories.NewPostgressBorrowRepository(database)
	
	bookService := services.Newbookservice(bookRepo)
	authService := services.NewAuthService(userRepo)
	borrowService:= services.NewBorrowService(database, bookRepo, borrowRepo)

	homeHandler := handlers.NewHomeHandler()
	bookHandler := handlers.NewBookHandler(bookService)
	authHandler := handlers.NewAuthHandler(authService)
	borrowHandler := handlers.NewBorrowHandler(borrowService)

	srv := server.New(homeHandler, bookHandler, authHandler, borrowHandler)
	
	log.Println("starting library api...")

	if err := srv.Start() ; err != nil {
		log.Fatal(err)
	}
}