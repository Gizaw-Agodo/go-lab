package handlers

import (
	"net/http"

	"github.com/gizaw/09-library-api/internal/dto"
	"github.com/gizaw/09-library-api/internal/middleware"
	"github.com/gizaw/09-library-api/internal/response"
	"github.com/gizaw/09-library-api/internal/services"
	"github.com/gizaw/09-library-api/internal/validation"
)

type BorrowHandler struct {
	service *services.BorrowService
}

func NewBorrowHandler(service *services.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		service: service,
	}
}

// BorrowBook godoc
//
//	@Summary		Borrow a book
//	@Description	Borrow an available book
//	@Tags			Borrow
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.BorrowRequest	true	"Borrow Request"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Failure		409		{object}	response.ErrorResponse
//	@Router			/borrow [post]
func (h *BorrowHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	req, err := dto.NewBorrowRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	claim, ok := middleware.GetUser(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized user")
		return
	}

	err = h.service.Borrow(r.Context(), claim.UserID, req.BookID)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.Created(w, response.SuccessResponse{
		Message: "book borrowed successfully",
	})
}

// ReturnBook godoc
//
//	@Summary		Return a book
//	@Description	Return a previously borrowed book
//	@Tags			Borrow
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.BorrowRequest	true	"Return Request"
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/return [post]
func (h *BorrowHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	req, err := dto.NewBorrowRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	claim, ok := middleware.GetUser(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.service.ReturnBook(r.Context(), claim.UserID, req.BookID); err != nil {
		response.DomainError(w, err)
		return
	}

	response.OK(w, response.SuccessResponse{
		Message: "book returned successfully",
	})
}

// ListBorrowedBooks godoc
//
//	@Summary		List my borrowed books
//	@Description	Get all books currently borrowed by the authenticated user
//	@Tags			Borrow
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	nil
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/borrow/me [get]
func (h *BorrowHandler) ListBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	claim, ok := middleware.GetUser(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	resp, err := h.service.ListBorrowedBooks(r.Context(), claim.UserID)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.OK(w, resp)
}