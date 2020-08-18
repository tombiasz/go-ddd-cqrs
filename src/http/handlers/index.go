package handlers

import (
	"encoding/json"
	"net/http"
)

type IndexResponse struct {
	Msg string `json:"msg"`
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	var response = &IndexResponse{"Hello World!"}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
