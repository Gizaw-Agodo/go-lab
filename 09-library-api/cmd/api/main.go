package main

import (
	"log"

	"github.com/gizaw/09-library-api/internal/config"
	"github.com/gizaw/09-library-api/internal/database"
	"github.com/gizaw/09-library-api/internal/handlers"
	"github.com/gizaw/09-library-api/internal/repositories"
	"github.com/gizaw/09-library-api/internal/server"
	"github.com/gizaw/09-library-api/internal/services"

	_ "github.com/gizaw/09-library-api/docs"
	"github.com/joho/godotenv"
)

// @title Library Management API
// @version 1.0
// @description REST API for Library Management.
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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