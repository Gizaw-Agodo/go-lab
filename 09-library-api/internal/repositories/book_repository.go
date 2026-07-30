package repositories

import (
	"context"
	"go-lab/09-library-api/internal/models"
)

type BookRepository interface {
	GetAll(ctx context.Context)([]models.Book, error)
	GetByID(ctx context.Context, id int) (*models.Book, error)
	Create(ctx context.Context, book *models.Book) (*models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
}
