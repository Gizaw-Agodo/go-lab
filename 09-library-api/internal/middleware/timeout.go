package middleware

import (
	"net/http"
)

func Timeout(next http.Handler) http.Handler {
	var handler http.HandlerFunc
	handler = func(w http.ResponseWriter, r *http.Request){
	}

	return handler
}