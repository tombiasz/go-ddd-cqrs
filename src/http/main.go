package main

import (
	"go-coupons/src/http/handlers"
	"go-coupons/src/http/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.ContentTypeJson)

	r.Route("/api/v1", func(r chi.Router) {

		r.MethodFunc("GET", "/", handlers.IndexHandler)
		r.MethodFunc("GET", "/coupons", handlers.GetCouponsHandler)
		r.MethodFunc("GET", "/coupons/{couponId}", handlers.GetCouponByIdHandler)
		r.MethodFunc("POST", "/coupons", handlers.RegisterCouponHandler)
		r.MethodFunc("POST", "/coupons/expire", handlers.ExpireCouponsHandler)
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
