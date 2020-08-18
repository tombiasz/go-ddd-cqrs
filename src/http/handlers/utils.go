package handlers

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Err string `json:"error"`
}

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)

	response := &jsonError{err.Error()}

	json.NewEncoder(w).Encode(response)
}
