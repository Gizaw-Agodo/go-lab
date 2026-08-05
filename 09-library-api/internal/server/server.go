package server

import (
	"net/http"

	"github.com/gizaw/09-library-api/internal/handlers"
)

type Server struct {
	httpServer *http.Server
	homeHandler *handlers.HomeHandler
	bookHandler *handlers.BookHandler
	authHanlder *handlers.AuthHandler
	borrowHandler *handlers.BorrowHandler
}

func New(
	homeHandler  *handlers.HomeHandler, 
	bookHandler *handlers.BookHandler,
	authHanlder *handlers.AuthHandler,
	borrowHandler *handlers.BorrowHandler,
) *Server {
	s := &Server{
		homeHandler: homeHandler,
		bookHandler: bookHandler,
		authHanlder: authHanlder,
		borrowHandler: borrowHandler,
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