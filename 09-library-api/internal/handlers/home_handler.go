package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Library API 🚀")
}