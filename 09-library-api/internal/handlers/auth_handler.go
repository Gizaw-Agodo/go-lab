package handlers

import (
	"go-lab/09-library-api/internal/dto"
	"go-lab/09-library-api/internal/response"
	"go-lab/09-library-api/internal/services"
	"go-lab/09-library-api/internal/validation"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService  *services.AuthService) *AuthHandler {
	return & AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request){
	req, err := dto.NewRegisterRequest(r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := validation.Validate(req); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return 
	}
	user, err := h.authService.Register(r.Context(), req)

	if err != nil {
		response.WriteError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, dto.NewRegisterResponse(user))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	req, err := dto.NewLoginRequest(r)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return 
	}
	user, token, err := h.authService.Login(r.Context(), req)
	if err != nil {
		response.WriteError(w, err)
		return
	}
	resp := dto.ToLoginResponse(user, token)
	response.JSON(w, http.StatusOK, resp)

}