package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errMsg string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   errMsg,
		Details: details,
	})
	if err != nil {
		return
	}
}
