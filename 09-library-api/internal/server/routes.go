package server

import (
	"go-lab/09-library-api/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recovery)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// home
	r.Get("/", s.homeHandler.Home)
	r.Route("/auth", func (r chi.Router) {
		r.Post("/register", s.authHanlder.RegisterUser)
		r.Post("/login", s.authHanlder.Login)
	})


	// books 
	r.Route("/books" , func(r chi.Router){
		
		r.Group(func (r chi.Router){
			r.Use(middleware.Authenticate)
			r.Get("/",s.bookHandler.GetBooks)
			r.Post("/", s.bookHandler.CreateBook)
			r.Get("/{id}", s.bookHandler.GetBook)
			r.Put("/{id}", s.bookHandler.UpdateBook)
			r.Delete("/{id}", s.bookHandler.DeleteBook)

		})

	})

	// borrow 
	r.Route("/borrow", func(r chi.Router){
		r.Post("/", s.borrowHandler.BorrowBook)
	})

	return r
}