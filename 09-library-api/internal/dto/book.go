package dto

import (
	"encoding/json"
	"errors"
	"go-lab/09-library-api/internal/models"
	"go-lab/09-library-api/internal/repositories"
	"net/http"
	"strconv"
)

type CreateBookRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}

type ListBooksRequest struct {
	Page int 
	Limit int 
}

func (r ListBooksRequest) ToRepositoryParams() repositories.ListBooksParams {
	return repositories.ListBooksParams{
		Page: r.Page,
		Limit: r.Limit,
		Offset: (r.Page - 1) * r.Limit,
	}
}

func (r CreateBookRequest) ToBook() *models.Book {
	return &models.Book{
		Title: r.Title,
		Author: r.Author,
	}
}

func NewListBooksRequest(r *http.Request)(*ListBooksRequest, error){
	params := ListBooksRequest{
		Page: 1,
		Limit: 10,
	}
	
	if p := r.URL.Query().Get("page"); p != "" {
		value, err := strconv.Atoi(p)
		if err != nil || value < 1 {
			return nil, errors.New("page must be greater than zero ")
		}
		params.Page = value
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		value, err := strconv.Atoi(l)
		if err != nil || value < 1 || value > 100 {
			return nil, errors.New("limit must be between 1 and 100") 
		}
		params.Limit = value
	}
	
	return &params, nil
}

func NewCreateBookRequest(r *http.Request)(*CreateBookRequest, error ){
	req:= &CreateBookRequest{} 
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUpdateBookRequest(r *http.Request)(*UpdateBookRequest, error ){
	req:= &UpdateBookRequest{} 
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}
	return req, nil
}