package main

import (
	"go-coupons/src/http/handlers"
	"go-coupons/src/http/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.ContentTypeJson)

	dbUrl := os.Getenv("DATABASE_URL")

	r.Route("/api/v1", func(r chi.Router) {

		r.MethodFunc("GET", "/", handlers.IndexHandler)
		r.MethodFunc("GET", "/coupons", handlers.CreateGetCouponsHandler(dbUrl))
		r.MethodFunc("GET", "/coupons/{couponId}", handlers.CreateGetCouponByIdHandler(dbUrl))
		r.MethodFunc("POST", "/coupons", handlers.CreateRegisterCouponHandler(dbUrl))
		r.MethodFunc("POST", "/coupons/expire", handlers.CreateExpireCouponsHandler(dbUrl))
		r.MethodFunc("POST", "/coupons/use", handlers.CreateUseCouponHandler(dbUrl))
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
