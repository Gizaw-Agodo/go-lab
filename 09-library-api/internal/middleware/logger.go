package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	var handler http.HandlerFunc
	handler = func(w http.ResponseWriter, r *http.Request){
		id := GetRequestId(r.Context())
		start := time.Now()
		log.Printf("Started %s %s %s",id, r.Method, r.URL)
		
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.Pattern, time.Since(start))

	}
	return handler
}