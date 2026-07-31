package repositories

import "go-lab/09-library-api/internal/models"

type ListBooksResult struct {
	Books []models.Book
	Total int
}