package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func createResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Error : %v", err)
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func createErrorResponse(w http.ResponseWriter, status int, response ErrorResponse) {
	createResponse(w, status, response)
}
