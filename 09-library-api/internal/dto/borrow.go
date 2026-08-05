package dto

import (
	"encoding/json"
	"net/http"
)

type BorrowRequest struct {
    BookID int64 `json:"book_id" validate:"required,min=1"`
}

type ListBorrowResponse struct {
	BookID int64 `json:"book_id"`
	Title string `json:"title"`
	Author string `json:"author"`
	BorrowedAt string `json:"borrowed_at"`
}

func NewBorrowRequest(r *http.Request) (*BorrowRequest, error) {
    var borrowRequest BorrowRequest

    if err := json.NewDecoder(r.Body).Decode(&borrowRequest); err != nil {
        return nil, err
    }

    return &borrowRequest, nil
}

