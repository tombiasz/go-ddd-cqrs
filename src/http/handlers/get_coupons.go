package handlers

import (
	"go-coupons/src/app/coupons/db"
	"go-coupons/src/app/coupons/queries"
	"net/http"
	"time"
)

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

func CreateGetCouponsHandler(dbUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		handler := &queries.GetCouponsQueryHandler{
			DbConnectionFactory: db.NewDbConnectionFactory(dbUrl),
		}

		result := handler.Query()

		var response []GetCouponsResponse
		for _, r := range result {
			response = append(response, GetCouponsResponse(r))
		}

		JSONResponse(w, response)
	}
}
