package server

import (
	"go-lab/09-library-api/internal/domain"
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
		admin := r.With( middleware.Authenticate, middleware.RequireRole(domain.RoleAdmin),)
		admin.Get("/",s.bookHandler.GetBooks)
		
		r.Group(func (r chi.Router){
			r.Use(middleware.Authenticate)
			r.Post("/", s.bookHandler.CreateBook)
			r.Get("/{id}", s.bookHandler.GetBook)
			r.Put("/{id}", s.bookHandler.UpdateBook)
			r.Delete("/{id}", s.bookHandler.DeleteBook)

		})

	})

	// borrow 
	r.Route("/borrows", func(r chi.Router){
		r.Use(middleware.Authenticate)
		r.Post("/", s.borrowHandler.BorrowBook)
		r.Post("/return", s.borrowHandler.ReturnBook)
		r.Get("/me", s.borrowHandler.ListBorrowedBooks)
	})

	return r
}