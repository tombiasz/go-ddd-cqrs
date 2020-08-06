package main

import (
	"encoding/json"
	"go-coupons/src/http/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type IndexResponse struct {
	Msg string `json:"msg"`
}

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	var response = &IndexResponse{"Hello World!"}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.ContentTypeJson)

	r.Route("/api/v1", func(r chi.Router) {

		r.MethodFunc("GET", "/", HandleIndex)
	})

	return r
}

func main() {
	appRouter := NewRouter()

	log.Printf("Starting server :8000")

	s := &http.Server{
		Addr:    ":8000",
		Handler: appRouter,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
