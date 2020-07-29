package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
)

type UseCouponCommand struct {
	Email string
	Code  string
}

type UseCouponCommandHandler struct {
	TimeProvider domain.TimeProvider
	Repository   coupon.Repository
}

func (h *UseCouponCommandHandler) Execute(cmd *UseCouponCommand) domain.DomainErrors {
	email, emailErr := coupon.NewEmail(cmd.Email)
	code, codeErr := coupon.NewCode(cmd.Code)

	err := domain.CombineDomainErrors(emailErr, codeErr)

	if err != nil {
		return err
	}

	coupon, repoErr := h.Repository.GetCouponByEmailAndCode(email, code)

	if repoErr != nil {
		return repoErr.AsDomainErrors()
	}

	domainErr := coupon.Use(h.TimeProvider)

	if domainErr != nil {
		return domainErr.AsDomainErrors()
	}

	repoErr = h.Repository.Save(coupon)

	if repoErr != nil {
		return repoErr.AsDomainErrors()
	}

	return nil
}
