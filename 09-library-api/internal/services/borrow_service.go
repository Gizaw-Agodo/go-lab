package services

import (
	"context"
	"database/sql"

	"github.com/gizaw/09-library-api/internal/domain"
	"github.com/gizaw/09-library-api/internal/dto"
	"github.com/gizaw/09-library-api/internal/models"
	"github.com/gizaw/09-library-api/internal/repositories"
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

func (s *BorrowService) ReturnBook(ctx context.Context, userID, bookID int64)error {
	
	// 1. initiate transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err 
	}
	defer tx.Rollback()

	// 2. create repositories 
	bookRepo := repositories.NewPostgressBookRepository(tx)
	borrowRepo := repositories.NewPostgressBorrowRepository(tx)

	// 3. check if there is active borrow 
	activeBorrow, err := borrowRepo.GetActiveBorrow(ctx, userID, bookID)
	if err != nil {
		return err 
	}

	// 4. update borrow 
	if err := borrowRepo.Return(ctx, activeBorrow.ID); err != nil {
		return err
	}

	// 5. update book 
	if err := bookRepo.UpdateAvailability(ctx, bookID, true); err != nil {
		return err 
	}

	// 6 commit 
	return tx.Commit()
}

func (s *BorrowService) ListBorrowedBooks(ctx context.Context, userID int64)([]dto.ListBorrowResponse, error){
	return s.borrowRepo.ListBorrowedBooks(ctx, userID)
}