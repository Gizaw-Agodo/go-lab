package handlers

import (
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/response"
	"go-lab/09-library-api/internal/services"
	"go-lab/09-library-api/internal/validation"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{
		service: service, 
	}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid book id")
		return
	}
	book,err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.DomainError(w, err)
		return 
	}
	response.OK(w, book)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll(r.Context())
	if err != nil {
		response.DomainError(w,err)
		return
	}
	response.OK(w, books)
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	
	params, err := dto.NewListBooksRequest(r)

	if err != nil {
		response.Error(w, http.StatusBadRequest,err.Error())
	}
	resp, err := h.service.ListBooks(r.Context(),params.ToRepositoryParams())
	if err != nil {
		response.DomainError(w, err)
		return
	}
	response.OK(w, resp)
}


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

	newBook, err := h.service.Create(r.Context(), req); 
	if err != nil {
		response.DomainError(w, err)
		return 
	}
	response.Created(w, newBook)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w,http.StatusBadRequest, "invalid book id")
		return 
	}
	
	req, err := dto.NewUpdateBookRequest(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest, err.Error())
		return 
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return 
	}

	err = h.service.Update(r.Context(),req, id)
	if err != nil {
		response.DomainError(w, err)
		return 
	}
	response.OK(w,req)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
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