package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter , status int , data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w,"Failed to encode json response", http.StatusInternalServerError)
	}
}