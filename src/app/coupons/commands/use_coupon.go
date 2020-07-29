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
		return domain.CombineDomainErrors(repoErr)
	}

	domainErr := coupon.Use(h.TimeProvider)

	if domainErr != nil {
		return domain.CombineDomainErrors(domainErr)
	}

	repoErr = h.Repository.Save(coupon)

	if repoErr != nil {
		// TODO: single err as array? blah
		return domain.CombineDomainErrors(repoErr)
	}

	return nil
}
