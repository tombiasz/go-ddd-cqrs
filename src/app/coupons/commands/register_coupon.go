package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
)

type RegisterCouponCommand struct {
	Email       string
	Description string
}

type RegisterCouponCommandHandler struct {
	TimeProvider domain.TimeProvider
	Repository   coupon.Repository
}

func (h *RegisterCouponCommandHandler) Execute(cmd *RegisterCouponCommand) domain.DomainErrors {
	coupon, domainErr := coupon.RegisterCoupon(
		cmd.Email,
		cmd.Description,
		coupon.DefaultCouponExpirationDays,
		h.TimeProvider,
	)

	if domainErr != nil {
		return domainErr
	}

	dbErr := h.Repository.Save(coupon)

	if dbErr != nil {
		// TODO: single err as array? blah
		return domain.CombineDomainErrors(dbErr)
	}

	return nil
}
