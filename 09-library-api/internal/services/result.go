package services

import "go-lab/09-library-api/internal/models"

type ListBooksResult struct {
	Books []models.Book
	Page  int
	Limit int
	Total int
}