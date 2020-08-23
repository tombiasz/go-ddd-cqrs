package handlers

import (
	"go-coupons/src/app/coupons/commands"
	"go-coupons/src/app/coupons/db"
	timeutils "go-coupons/src/utils/time"
	"net/http"
)

func CreateExpireCouponsHandler(dbUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := db.NewCouponRepository(
			db.NewDbConnectionFactory(dbUrl),
		)

		handler := &commands.ExpireCouponsCommandHandler{
			Repository:   repo,
			TimeProvider: &timeutils.RealTimeProvider{},
		}

		errDomain := handler.Execute()

		if errDomain != nil {
			JSONDomainErrors(w, errDomain, 400)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
