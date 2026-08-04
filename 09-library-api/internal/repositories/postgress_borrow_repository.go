package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-lab/09-library-api/internal/database"
	"go-lab/09-library-api/internal/domain"
	"go-lab/09-library-api/internal/models"
)

type PostgressBorrowRepository struct {
	db database.DBTX
}

func NewPostgressBorrowRepository(db database.DBTX)*PostgressBorrowRepository {
	return &PostgressBorrowRepository{
		db : db,
	}
}

func (r *PostgressBorrowRepository)Create(ctx context.Context, req *models.Borrow)(*models.Borrow, error){
	borrow := &models.Borrow{
		UserID: req.UserID,
		BookID: req.BookID,
	}

	query := `
		INSERT INTO borrows (
			user_id,
			book_id
		)
		VALUES ($1, $2)
		RETURNING
			id,
			borrowed_at,
			returned_at`
		
		if err := r.db.QueryRowContext(ctx, query, req.UserID, req.BookID).Scan(
			&borrow.ID,
			&borrow.BorrowedAt,
			borrow.ReturnedAt,
		); err != nil {
			return nil, err
		}

		return borrow, nil 
}

func (r *PostgressBorrowRepository) GetActiveBorrow(ctx context.Context, userID int64 , bookID int64)(*models.Borrow, error){
	borrow := &models.Borrow{

	}
	query := `
		SELECT
			id,
			user_id,
			book_id,
			borrowed_at,
			returned_at
		FROM borrows
		WHERE
			user_id = $1
			AND book_id = $2
			AND returned_at IS NULL`
		
	if err := r.db.QueryRowContext(ctx, query, userID, bookID).Scan(
		&borrow.ID, 
		&borrow.UserID, 
		&borrow.BookID,
		&borrow.BorrowedAt, 
		&borrow.ReturnedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBorrowNotFound
		}
		return nil, err
	}

	return borrow, nil 

}

func (r *PostgressBorrowRepository)Return(ctx context.Context, borrowID int64) error {
	query := `
		UPDATE borrows
		SET returned_at = NOW()
		WHERE id = $1`
	
	result, err := r.db.ExecContext(
		ctx,
		query,
		borrowID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrBorrowNotFound
	}

	return nil
}

func (r *PostgressBorrowRepository) ListByUser( ctx context.Context, userID int64,) ([]models.Borrow, error) {
	query := `
		SELECT
			id,
			user_id,
			book_id,
			borrowed_at,
			returned_at
		FROM borrows
		WHERE user_id = $1
		ORDER BY borrowed_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var borrows []models.Borrow

	for rows.Next(){
		var borrow models.Borrow
		if err := rows.Scan(
			&borrow.ID, 
			&borrow.UserID, 
			&borrow.BookID, 
			&borrow.BorrowedAt, 
			&borrow.ReturnedAt,
		); err != nil {
			return nil, err
		}
		
		borrows = append(borrows, borrow)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return borrows, nil 

}


func (r *PostgressBorrowRepository)CountActiveBorrow(ctx context.Context, userID int64)(int, error){
	query := `
		SELECT COUNT(*)
		FROM borrows
		WHERE
			user_id = $1
			AND returned_at IS NULL
	`

	var count int

	if err := r.db.QueryRowContext(ctx, query, userID,).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}