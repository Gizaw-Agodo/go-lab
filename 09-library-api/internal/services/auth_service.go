package services

import (
	"context"
	"errors"

	"github.com/gizaw/09-library-api/internal/auth"
	"github.com/gizaw/09-library-api/internal/domain"
	"github.com/gizaw/09-library-api/internal/dto"
	"github.com/gizaw/09-library-api/internal/models"
	"github.com/gizaw/09-library-api/internal/repositories"
)

type AuthService struct {
	userRepo *repositories.PostgressUserRepository
}

func NewAuthService(repo *repositories.PostgressUserRepository) *AuthService{
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest)(*models.User, error){

	pasword_hash, err := auth.Hash(req.Password)
	if err != nil {
		return nil, err 
	}
	
	params := repositories.CreateUserParams{
		Email: req.Email,
		Name: req.Name,
		PasswordHash: pasword_hash,
	}
	return s.userRepo.Create(ctx, params)
}

func (s *AuthService)Login(ctx context.Context, req *dto.LoginRequest)(*models.User,*string, error){
	user, err := s.userRepo.GetByEmail(ctx,req.Email )
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil,nil, domain.ErrInvalidCredentials
		}
		return nil,nil, err
	}

	if err := auth.Compare(user.PasswordHash, req.Password); err != nil {
		return nil,nil, domain.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken( user.ID, user.Email,user.Role)
	
	if err != nil {
		return nil,nil, err
	}
	return user, &token, nil

}