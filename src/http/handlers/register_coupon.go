package handlers

import (
	"encoding/json"
	"go-coupons/src/app/coupons/commands"
	"go-coupons/src/app/coupons/db"
	timeutils "go-coupons/src/utils/time"
	"net/http"
)

type RegisterCouponRequest struct {
	Email       *string `json:"email`
	Description *string `json:"description"`
}

type RegisterCouponResponse struct {
	Code string `json:"code`
}

func CreateRegisterCouponHandler(dbUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var req RegisterCouponRequest

		err := d.Decode(&req)

		if err != nil {
			// bad JSON or unrecognized json field
			// http.Error(w, err.Error(), http.StatusBadRequest)
			JSONError(w, err, http.StatusBadRequest)
			return
		}

		repo := db.NewCouponRepository(
			db.NewDbConnectionFactory(dbUrl),
		)

		handler := &commands.RegisterCouponCommandHandler{
			Repository:   repo,
			TimeProvider: &timeutils.RealTimeProvider{},
		}

		cmd := &commands.RegisterCouponCommand{
			Email:       *req.Email,
			Description: *req.Description,
		}

		result, errDomain := handler.Execute(cmd)

		if errDomain != nil {
			JSONDomainErrors(w, errDomain, 400)
			return
		}

		response := RegisterCouponResponse(*result)

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
