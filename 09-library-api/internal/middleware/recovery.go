package middleware

import (
	"log"
	"net/http"

	"github.com/gizaw/09-library-api/internal/response"
)

func Recovery(next http.Handler) http.Handler {
	var handler http.HandlerFunc
	handler = func (w http.ResponseWriter, r *http.Request) {
		defer func(){
			if err := recover(); err != nil {
				log.Printf("pannic recovered %v", err)
				response.Error( w, http.StatusInternalServerError, "internal server error")
			}
		}()
		
		next.ServeHTTP(w, r)
	}

	return handler
}