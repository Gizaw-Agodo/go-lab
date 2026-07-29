package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	// home
	r.Get("/", s.homeHandler.Home)

	// books 
	r.Route("/books" , func(r chi.Router){
		r.Get("/",s.bookHandler.GetBooks)
		r.Post("/", s.bookHandler.CreateBook)
		r.Get("/{id}", s.bookHandler.GetBook)
		r.Put("/{id}", s.bookHandler.UpdateBook)
		r.Delete("/{id}", s.bookHandler.DeleteBook)

	})

	return r
}