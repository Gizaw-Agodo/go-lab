package middleware

import (
	"context"
	"go-lab/09-library-api/internal/auth"
	"go-lab/09-library-api/internal/response"
	"net/http"
	"strings"
)

const UserContextKey ContextKey = "user"

func Authenticate(next http.Handler) http.Handler {
	var handler http.HandlerFunc

	handler = func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			response.Error(w, http.StatusUnauthorized, "Missing authorization header")
			return
		}

		token, ok := strings.CutPrefix(authHeader, "Bearer ")
		
		if !ok {
			response.Error(w, http.StatusUnauthorized, "Missing authorization header")
			return
		}

		claim, err := auth.ParseToken(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(),UserContextKey, claim)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return handler
}

func GetUser(ctx context.Context)(*auth.Claims, bool){
	claim, ok := ctx.Value(UserContextKey).(*auth.Claims)
	if !ok {
		return nil, false
	}
	return claim, true
}