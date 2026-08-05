package repositories

import (
	"context"
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/models"
)

type BorrowRepository interface {
	Create(cxt context.Context, borrow *models.Borrow)(*models.Borrow, error)
	GetActiveBorrow(ctx context.Context, userID int64 , bookID int64)(*models.Borrow, error)
	Return(ctx context.Context, borrowID int64) error
	ListByUser(ctx context.Context, userID int64)([]models.Borrow, error)
	CountActiveBorrow(ctx context.Context, userID int64)(int, error)
	ListBorrowedBooks(ctx context.Context, userID int64)([]dto.ListBorrowResponse, error)
}