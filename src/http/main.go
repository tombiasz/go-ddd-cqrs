package main

import (
	"encoding/json"
	"go-coupons/src/http/middlewares"
	"log"
	"net/http"
	"os"
	"time"

	"go-coupons/src/app/coupons/db"
	"go-coupons/src/app/coupons/queries"

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

type GetCouponsResponse struct {
	Id          string     `json:"id"`
	Code        string     `json:"code"`
	Email       string     `json:"email"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Expdays     uint8      `json:"expiriesInDays"`
	ActivatedAt time.Time  `json:"activatedAt"`
	ExpiredAt   *time.Time `json:"expiredAt"`
	UsedAt      *time.Time `json:"usedAt"`
}

func GetCouponsHandler(w http.ResponseWriter, _ *http.Request) {
	handler := &queries.GetCouponsQueryHandler{
		DbConnectionFactory: db.NewDbConnectionFactory(os.Getenv("DATABASE_URL")),
	}

	result := handler.Query()

	var response []GetCouponsResponse
	for _, r := range result {
		response = append(response, GetCouponsResponse(r))
	}

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
		r.MethodFunc("GET", "/coupons", GetCouponsHandler)
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
