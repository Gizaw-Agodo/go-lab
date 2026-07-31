package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-lab/09-library-api/internal/models"
)

type PostgressBookRepository struct {
	db *sql.DB
}

func NewPostgressBookRepository(db *sql.DB) *PostgressBookRepository {
	return &PostgressBookRepository{
		db: db,
	}
}

func (r *PostgressBookRepository) GetAll(ctx context.Context)([]models.Book, error){
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, title, author, created_at, updated_at
		FROM books
		ORDER BY id 
	`)

	if err != nil {
		return nil , err 
	}
	defer rows.Close()
	var books []models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(
			&book.ID, 
			&book.Title, 
			&book.Author,
			&book.CreatedAt, 
			&book.UpdatedAt); err != nil {
			return nil, err 
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil , err
	}

	return books, nil 
}


func(r *PostgressBookRepository) List(ctx context.Context, params ListBooksParams) (*ListBooksResult, error){
	query := `
		SELECT
			id,
			title,
			author,
			created_at,
			updated_at
		FROM books
		ORDER BY id
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	

	books := make([]models.Book,0, params.Limit)

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(
			&book.ID, 
			&book.Title, 
			&book.Author, 
			&book.CreatedAt, 
			&book.UpdatedAt,
		); err != nil {
			return nil , err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	const countQuery = `
		SELECT COUNT(*)
		FROM books
	`

	var total int

	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, err
	}

	return &ListBooksResult{ Books: books, Total: total, }, nil
}

func (r *PostgressBookRepository) GetByID(ctx context.Context, id int) (*models.Book, error){
	query := `
		SELECT id, title, author, created_at, updated_at
		FROM books
		WHERE id = $1
	`
	var book models.Book

	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&book.ID, 
		&book.Title, 
		&book.Author,
		&book.CreatedAt, 
		&book.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &book, nil
}

func (r *PostgressBookRepository) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	query := `
		INSERT INTO books (title, author)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	if err := r.db.QueryRowContext(ctx, query, book.Title, book.Author).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.UpdatedAt,
	); err != nil {
		return nil, err 
	}


	return book, nil 
}

func (r *PostgressBookRepository) Update(ctx context.Context, book * models.Book) error {
	query := `
		UPDATE books
		SET
			title = $1,
			author = $2,
			updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at
	`

	err := r.db.QueryRowContext(ctx, query, book.Title, book.Author, book.ID).Scan(
		&book.UpdatedAt,
	)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err){
			return sql.ErrNoRows
		}
		return err 
	}
	return nil 
}

func (r *PostgressBookRepository)Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM books
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err 
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err 
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil 
}