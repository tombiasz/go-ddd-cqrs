package handlers

import (
	"encoding/json"
	"go-coupons/src/app/coupons/domain"
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

type jsonDomainError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func JSONDomainErrors(w http.ResponseWriter, errs domain.DomainErrors, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)

	var response []*jsonDomainError

	for c, m := range errs.AsMap() {
		response = append(response, &jsonDomainError{Code: c, Message: m})
	}

	json.NewEncoder(w).Encode(response)
}

func JSONResponse(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
