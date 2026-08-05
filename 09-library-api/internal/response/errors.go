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

func DomainError(w http.ResponseWriter, err error) {
	switch {
		case errors.Is(err, domain.ErrBookNotFound):
			writeJSON(w, http.StatusNotFound , ErrorResponse{Error: err.Error()})
		case errors.Is(err, domain.ErrDuplicateEmail): 
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		
		// borrow 
		case errors.Is(err, domain.ErrBorrowNotFound): 
			writeJSON(w, http.StatusNotFound, ErrorResponse{Error: err.Error()})
		case errors.Is(err, domain.ErrBorrowLimitReached): 
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		case errors.Is(err, domain.ErrBookUnavailable): 
			writeJSON(w, http.StatusNotFound, ErrorResponse{Error: err.Error()})

		default:
			log.Printf("unexpected err : %v", err)
			writeJSON(w, http.StatusInternalServerError,ErrorResponse{Error: "Internal server error"})
	}
}


func Error(w http.ResponseWriter, statuscode int,  errMessage string){
	writeJSON(w, statuscode, ErrorResponse{Error:errMessage } )
}