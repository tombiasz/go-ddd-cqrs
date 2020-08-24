package main

import (
	"fmt"
	"go-coupons/src/http/config"
	"go-coupons/src/http/handlers"
	"go-coupons/src/http/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(conf config.AppConfig) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.ContentTypeJson)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.MethodFunc("GET", "/", handlers.IndexHandler)
		r.MethodFunc("GET", "/coupons", handlers.CreateGetCouponsHandler(conf.DbUrl))
		r.MethodFunc("GET", "/coupons/{couponId}", handlers.CreateGetCouponByIdHandler(conf.DbUrl))
		r.MethodFunc("POST", "/coupons", handlers.CreateRegisterCouponHandler(conf.DbUrl))
		r.MethodFunc("POST", "/coupons/expire", handlers.CreateExpireCouponsHandler(conf.DbUrl))
		r.MethodFunc("POST", "/coupons/use", handlers.CreateUseCouponHandler(conf.DbUrl))
	})

	return r
}

func main() {
	conf := config.New()
	appRouter := NewRouter(conf)

	log.Printf("Starting server :%s", conf.AppPort)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.AppPort),
		Handler: appRouter,
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in f %s", r)
		}
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
