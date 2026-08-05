package repositories

import "github.com/gizaw/09-library-api/internal/models"

type ListBooksResult struct {
	Books []models.Book
	Total int
}