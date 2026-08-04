package response

import "net/http"

type SuccessResponse struct {
	Message string `json:"message"` 
}


func OK(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusCreated, data)
}

func Accepted(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusAccepted, data)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}