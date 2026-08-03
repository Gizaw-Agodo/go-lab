package dto

import (
	"encoding/json"
	"go-lab/09-library-api/internal/models"
	"go-lab/09-library-api/internal/validation"
	"net/http"
	"time"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type RegisterResponse struct {
	ID int64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	Token *string       `json:"token"`
	User  RegisterResponse `json:"user"`
}

func NewRegisterRequest(r *http.Request) (*RegisterRequest, error){
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req) ; err != nil {
		return nil, err
	}
	return &req, nil
}

func NewRegisterResponse(user *models.User)RegisterResponse{
	return RegisterResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}


func NewLoginRequest(r *http.Request)(*LoginRequest, error){
	var req *LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	if err := validation.Validate(req); err != nil {
		return nil, err
	}

	return req, nil 
}

func ToLoginResponse(user *models.User, token *string) *LoginResponse {
	return &LoginResponse{
		Token: token,
		User: NewRegisterResponse(user),
	}
}