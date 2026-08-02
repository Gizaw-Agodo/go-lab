package services

import (
	"context"
	"go-lab/09-library-api/internal/auth"
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/models"
	"go-lab/09-library-api/internal/repositories"
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