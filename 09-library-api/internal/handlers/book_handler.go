package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/gizaw/09-library-api/internal/dto"
	"github.com/gizaw/09-library-api/internal/response"
	"github.com/gizaw/09-library-api/internal/services"
	"github.com/gizaw/09-library-api/internal/validation"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

// GetBook godoc
//
//	@Summary		Get book by ID
//	@Description	Get a single book by its ID
//	@Tags			Books
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	dto.BookResponse
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/books/{id} [get]
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid book id")
		return
	}

	book, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.OK(w, book)
}

// GetBooks godoc
//
//	@Summary		List books
//	@Description	List books with pagination
//	@Tags			Books
//	@Produce		json
//	@Param			page	query		int	false	"Page number"
//	@Param			limit	query		int	false	"Page size"
//	@Success		200		{object}	dto.ListBooksResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/books [get]
func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	params, err := dto.NewListBooksRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.service.ListBooks(r.Context(), params)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.OK(w, resp)
}

// CreateBook godoc
//
//	@Summary		Create book
//	@Description	Create a new book
//	@Tags			Books
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateBookRequest	true	"Create Book"
//	@Success		201		{object}	dto.BookResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	req, err := dto.NewCreateBookRequest(r)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	newBook, err := h.service.Create(r.Context(), req)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.Created(w, newBook)
}

// UpdateBook godoc
//
//	@Summary		Update book
//	@Description	Update an existing book
//	@Tags			Books
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Book ID"
//	@Param			request	body		dto.UpdateBookRequest		true	"Update Book"
//	@Success		200		{object}	dto.BookResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/books/{id} [put]
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid book id")
		return
	}

	req, err := dto.NewUpdateBookRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Update(r.Context(), req, id)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.OK(w, req)
}

// DeleteBook godoc
//
//	@Summary		Delete book
//	@Description	Delete a book
//	@Tags			Books
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		204
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid book id")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		response.DomainError(w, err)
		return
	}

	response.NoContent(w)
}