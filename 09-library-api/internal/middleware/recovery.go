package middleware

import (
	"go-lab/09-library-api/internal/response"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	var handler http.HandlerFunc
	handler = func (w http.ResponseWriter, r *http.Request) {
		defer func(){
			if err := recover(); err != nil {
				log.Printf("pannic recovered %v", err)
				response.JSON(w,http.StatusInternalServerError, response.ErrorResponse{Error: "internal server error"})
			}
		}()
		
		next.ServeHTTP(w, r)
	}

	return handler
}