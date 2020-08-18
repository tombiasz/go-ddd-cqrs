package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
)

type RegisterCouponCommand struct {
	Email       string
	Description string
}

type RegisterCouponCommandResult struct {
	Code string
}

type RegisterCouponCommandHandler struct {
	TimeProvider domain.TimeProvider
	Repository   coupon.Repository
}

func (h *RegisterCouponCommandHandler) Execute(cmd *RegisterCouponCommand) (*RegisterCouponCommandResult, domain.DomainErrors) {
	coupon, domainErr := coupon.RegisterCoupon(
		cmd.Email,
		cmd.Description,
		coupon.DefaultCouponExpirationDays,
		h.TimeProvider,
	)

	if domainErr != nil {
		return nil, domainErr
	}

	dbErr := h.Repository.Save(coupon)

	if dbErr != nil {
		return nil, dbErr.AsDomainErrors()
	}

	return &RegisterCouponCommandResult{coupon.Code()}, nil
}
