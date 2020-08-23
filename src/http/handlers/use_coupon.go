package handlers

import (
	"encoding/json"
	"go-coupons/src/app/coupons/commands"
	"go-coupons/src/app/coupons/db"
	timeutils "go-coupons/src/utils/time"
	"net/http"
)

type UseCouponRequest struct {
	Email *string `json:"email`
	Code  *string `json:"code"`
}

func CreateUseCouponHandler(dbUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var req UseCouponRequest

		err := d.Decode(&req)

		if err != nil {
			JSONError(w, err, http.StatusBadRequest)
			return
		}

		repo := db.NewCouponRepository(
			db.NewDbConnectionFactory(dbUrl),
		)

		handler := &commands.UseCouponCommandHandler{
			Repository:   repo,
			TimeProvider: &timeutils.RealTimeProvider{},
		}

		cmd := &commands.UseCouponCommand{
			Email: *req.Email,
			Code:  *req.Code,
		}

		errDomain := handler.Execute(cmd)

		if errDomain != nil {
			JSONDomainErrors(w, errDomain, 400)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
