package server

import (
	"go-lab/09-library-api/internal/handlers"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	homeHandler *handlers.HomeHandler
	bookHandler *handlers.BookHandler
	authHanlder *handlers.AuthHandler
}

func New(
	homeHandler  *handlers.HomeHandler, 
	bookHandler *handlers.BookHandler,
	authHanlder *handlers.AuthHandler,
) *Server {
	s := &Server{
		homeHandler: homeHandler,
		bookHandler: bookHandler,
		authHanlder: authHanlder,
	}
	s.httpServer = &http.Server{
		Addr: ":8080",
		Handler: s.routes(),
	}
	return s
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}