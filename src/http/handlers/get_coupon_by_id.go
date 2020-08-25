package handlers

import (
	"errors"
	"go-coupons/src/app/coupons/db"
	"go-coupons/src/app/coupons/queries"
	"net/http"
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

func CreateGetCouponByIdHandler(dbUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		couponId := chi.URLParam(r, "couponId")

		handler := &queries.GetCouponByIdQueryHandler{
			DbConnectionFactory: db.NewDbConnectionFactory(dbUrl),
		}

		result := handler.Query(couponId)

		if result == nil {
			JSONError(w, errors.New("coupon not found"), 404)
			return
		}

		response := GetCouponByIdResponse(*result)

		JSONResponse(w, response)

	}
}
