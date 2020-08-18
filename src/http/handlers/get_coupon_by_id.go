package handlers

import (
	"encoding/json"
	"go-coupons/src/app/coupons/db"
	"go-coupons/src/app/coupons/queries"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

type GetCouponByIdResponse struct {
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

func GetCouponByIdHandler(w http.ResponseWriter, r *http.Request) {
	couponId := chi.URLParam(r, "couponId")

	handler := &queries.GetCouponByIdQueryHandler{
		DbConnectionFactory: db.NewDbConnectionFactory(os.Getenv("DATABASE_URL")),
	}

	result := handler.Query(couponId)

	response := GetCouponByIdResponse(result)

	// TODO:
	// - handle 404
	// handle other errors
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
