package services

import (
	"context"
	"database/sql"
	"go-lab/09-library-api/internal/domain"
	"go-lab/09-library-api/internal/models"
	"go-lab/09-library-api/internal/repositories"
)

type BorrowService struct {
	db *sql.DB
	bookRepo repositories.BookRepository
	borrowRepo repositories.BorrowRepository
}

func NewBorrowService(
	db *sql.DB,
	bookRepo repositories.BookRepository,
	borrowRepo repositories.BorrowRepository,
) *BorrowService {
	return &BorrowService{
		db:         db,
		bookRepo:   bookRepo,
		borrowRepo: borrowRepo,
	}
}

func(s *BorrowService) Borrow(ctx context.Context, userID int64, bookID int64)error{
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// transaction repositories

	bookRepo := repositories.NewPostgressBookRepository(tx)
	borrowRepo := repositories.NewPostgressBorrowRepository(tx)

	// book exists ?
	book, err := bookRepo.GetByID(ctx, int(bookID))
	if err != nil {
		return err
	}
	
	// book available 
	if !book.Available {
		return domain.ErrBookUnavailable
	}

	// count active borrow 
	count, err := borrowRepo.CountActiveBorrow(ctx, userID)
	if err!= nil {
		return err
	}
	if count >= 3 {
		return domain.ErrBorrowLimitReached
	}

	//create borrow 
	_, err = borrowRepo.Create(ctx, &models.Borrow{UserID: userID, BookID: bookID})
	if err != nil {
		return err 
	}

	// update availablity
	if err := bookRepo.UpdateAvailability(ctx, bookID, false); err != nil {
		return err
	}

	return tx.Commit()
}