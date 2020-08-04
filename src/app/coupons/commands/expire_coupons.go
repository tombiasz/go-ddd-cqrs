package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
)

type ExpireCouponsCommandHandler struct {
	TimeProvider domain.TimeProvider
	Repository   coupon.Repository
}

func (h *ExpireCouponsCommandHandler) Execute() domain.DomainErrors {
	coupons, repoErr := h.Repository.GetExpiredCoupons()

	if repoErr != nil {
		return repoErr.AsDomainErrors()
	}

	for _, coupon := range coupons {
		domainErr := coupon.Expire(h.TimeProvider)

		if domainErr != nil {
			return domainErr.AsDomainErrors()
		}

		repoErr = h.Repository.Save(coupon)

		if repoErr != nil {
			return repoErr.AsDomainErrors()
		}
	}

	return nil
}
