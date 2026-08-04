package handlers

import (
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/middleware"
	"go-lab/09-library-api/internal/response"
	"go-lab/09-library-api/internal/services"
	"go-lab/09-library-api/internal/validation"
	"net/http"
)

type BorrowHandler struct {
	service *services.BorrowService
}

func NewBorrowHandler(service *services.BorrowService)*BorrowHandler{
	return &BorrowHandler{
		service: service,
	}
}
func (h *BorrowHandler) BorrowBook(w http.ResponseWriter, r *http.Request){
	req, err := dto.NewBorrowRequest(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest, err.Error())
		return 
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w,http.StatusBadRequest, err.Error())
		return 
	}

	claim, ok  := middleware.GetUser(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized,"unauthorized user")
		return 
	}

	err = h.service.Borrow(r.Context(),claim.UserID,req.BookID )
	if err != nil {
		response.DomainError(w, err)
		return 
	}

	response.Created(w, response.SuccessResponse{Message: "book borrowed successfully"})

}