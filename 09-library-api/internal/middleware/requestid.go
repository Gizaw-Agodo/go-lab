package middleware

import (
	"context"
	"net/http"
	"github.com/google/uuid"
)

type ContextKey string 

var RequstIdKey ContextKey = "request_id"


func RequestID(next http.Handler) http.Handler {
	var handler http.HandlerFunc

	handler = func(w http.ResponseWriter, r *http.Request){
		requestId := uuid.NewString()
		ctx := context.WithValue(r.Context(),RequstIdKey, requestId )
		w.Header().Set("X-Request-ID", requestId)

		next.ServeHTTP(w,r.WithContext(ctx))
	}

	return handler
}

func GetRequestId(ctx context.Context)string{
	id,ok := ctx.Value(RequstIdKey).(string)
	if !ok  {
		return ""
	}
	return id
}