package handlers

import (
	"net/http"

	"github.com/gizaw/09-library-api/internal/dto"
	"github.com/gizaw/09-library-api/internal/response"
	"github.com/gizaw/09-library-api/internal/services"
	"github.com/gizaw/09-library-api/internal/validation"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterUser godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterRequest	true	"Register Request"
//	@Success		201		{object}	dto.RegisterResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		409		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/register [post]
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	req, err := dto.NewRegisterRequest(r)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := validation.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authService.Register(r.Context(), req)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	response.Created(w, dto.NewRegisterResponse(user))
}

// Login godoc
//
//	@Summary		Login
//	@Description	Authenticate a user and return a JWT
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login Request"
//	@Success		200		{object}	dto.LoginResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req, err := dto.NewLoginRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := h.authService.Login(r.Context(), req)
	if err != nil {
		response.DomainError(w, err)
		return
	}

	resp := dto.ToLoginResponse(user, token)
	response.OK(w, resp)
}