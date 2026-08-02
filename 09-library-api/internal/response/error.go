package response

import (
	"errors"
	"go-lab/09-library-api/internal/domain"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"` 
}

func WriteError(w http.ResponseWriter, err error) {
	switch {
		case errors.Is(err, domain.ErrBookNotFound):
			JSON(w, http.StatusNotFound , ErrorResponse{Error: err.Error()})
		case errors.Is(err, domain.ErrDuplicateEmail): 
			JSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		default:
			log.Printf("unexpected err : %v", err)
			JSON(w, http.StatusInternalServerError,ErrorResponse{Error: "Internal server error"})
	}
}