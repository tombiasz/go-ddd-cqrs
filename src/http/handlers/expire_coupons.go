package handlers

import (
	"go-coupons/src/app/coupons/commands"
	"go-coupons/src/app/coupons/db"
	timeutils "go-coupons/src/utils/time"
	"net/http"
	"os"
)

func ExpireCouponsHandler(w http.ResponseWriter, r *http.Request) {
	repo := db.NewCouponRepository(
		db.NewDbConnectionFactory(os.Getenv("DATABASE_URL")),
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
