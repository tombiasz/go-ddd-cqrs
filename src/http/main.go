package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write([]byte("Hello World!"))
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("GET", "/", HandleIndex)
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
