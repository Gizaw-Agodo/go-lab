package handlers

import (
	"encoding/json"
	"fmt"
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/models"
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
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "Invalid book id"})
		return
	}
	book,err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.WriteError(w, err)
		return 
	}
	response.JSON(w, http.StatusOK, book)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll(r.Context())
	if err != nil {
		response.WriteError(w,err)
		return
	}
	response.JSON(w, http.StatusOK, books)
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	
	params, err := dto.NewListBooksRequest(r)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
	}
	resp, err := h.service.ListBooks(r.Context(),params.ToRepositoryParams())
	if err != nil {
		response.WriteError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, resp)
}


func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBookRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := validation.Validate(req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	book := req.ToBook()
	
	newBook, err := h.service.Create(r.Context(), book); 
	
	if err != nil {
		response.WriteError(w, err)
		return 
	}
	response.JSON(w,http.StatusCreated, newBook)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateBookRequest
	var decoder = json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.JSON(w,http.StatusBadRequest, response.ErrorResponse{Error: "Invalid book id"})
		return 
	}
	
	if err := decoder.Decode(&req); err != nil {
		response.JSON(w,http.StatusBadRequest, response.ErrorResponse{Error: "Invalid request body"})
		return 
	}

	if err := validation.Validate(req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return 
	}

	book := &models.Book{
		ID: int(id),
		Title: req.Title,
		Author: req.Author,
	}
	
	err = h.service.Update(r.Context(),book)
	if err != nil {
		response.WriteError(w, err)
		return 
	}

	response.JSON(w, http.StatusOK, req)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete a book")
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: "invalid book id "})
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		response.WriteError(w, err)
		return 
	}
	w.WriteHeader(http.StatusNoContent)
}