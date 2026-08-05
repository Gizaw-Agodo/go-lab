package services

import (
	"context"
	"database/sql"
	"errors"
	"go-lab/09-library-api/internal/domain"
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/models"
	"go-lab/09-library-api/internal/repositories"
)

type BookService struct {
	repo repositories.BookRepository
}

func Newbookservice(repo repositories.BookRepository) *BookService {
	return &BookService{
		repo : repo,
	}
}

func (s *BookService) GetAll(ctx context.Context)([]models.Book, error){
	return s.repo.GetAll(ctx)
}

func (s *BookService) ListBooks(ctx context.Context,req *dto.ListBooksRequest ) (*ListBooksResult, error) {
	params := repositories.ListBooksParams{
		Page: req.Page,
		Limit: req.Limit,
		Offset: (req.Page - 1)* req.Limit,
	}
	
	result, err := s.repo.List(ctx, params)
	if err != nil {
		return nil, err 
	}
	return &ListBooksResult{
		Books: result.Books,
		Page: params.Page,
		Limit: params.Limit,
		Total: result.Total,

	}, nil
}

func (s *BookService) GetByID(ctx context.Context, id int) (*models.Book, error){
	book, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, err
	}
	return book, err
}

func (s *BookService) Create(ctx context.Context, req *dto.CreateBookRequest) (*models.Book, error) {
	book := &models.Book{
		Title: req.Title,
		Author: req.Author,
	}
	return s.repo.Create(ctx, book) 
}

func (s *BookService) Update(ctx context.Context, req *dto.UpdateBookRequest , bookId int) error {
	book := &models.Book{
		ID: bookId,
		Title: req.Title,
		Author: req.Author,
	}
	err := s.repo.Update(ctx, book)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return domain.ErrBookNotFound 
		}
		return err 
	}
	return err
}

func (s *BookService) Delete(ctx context.Context, id int ) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrBookNotFound
		}
		return err 
	}
	return err 

}