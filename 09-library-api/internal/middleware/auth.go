package middleware

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/gizaw/09-library-api/internal/auth"
	"github.com/gizaw/09-library-api/internal/response"
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

func RequireRole(roles ...string) func(http.Handler) http.Handler {

	return func (next http.Handler) http.Handler{
		var handler http.HandlerFunc
		handler = func(w http.ResponseWriter, r *http.Request){
			user, ok := GetUser(r.Context())
			if !ok {
				response.Error(w, http.StatusUnauthorized, "unauthorized user")
				return 
			}
			if slices.Contains(roles, user.Role) {
					next.ServeHTTP(w, r)
					return
			}

			response.Error(w, http.StatusForbidden, "Insuficcient permissions")
				
		}

		return handler

	}

	
}