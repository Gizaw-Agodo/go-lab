package dto

import (
	"encoding/json"
	"go-lab/09-library-api/internal/validation"
	"net/http"
)

type BorrowRequest struct {
    BookID int64 `json:"book_id" validate:"required,min=1"`
}

func NewBorrowRequest(r *http.Request)(*BorrowRequest, error){
	var borrowRequest *BorrowRequest
	err := json.NewDecoder(r.Body).Decode(borrowRequest)
	if err != nil {
		return nil, err 
	}

	if err := validation.Validate(borrowRequest); err != nil {
		return nil, err
	}
	return borrowRequest, nil
}