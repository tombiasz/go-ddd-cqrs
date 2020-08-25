package handlers

import (
	"net/http"
)

type IndexResponse struct {
	Msg string `json:"msg"`
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	var response = &IndexResponse{"coupons api v1"}

	JSONResponse(w, response)
}
