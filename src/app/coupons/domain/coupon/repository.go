package coupon

import "go-coupons/src/app/coupons/domain"

type Repository interface {
	Save(coupon *Coupon) *domain.DomainError

	GetCouponByEmailAndCode(*Email, *Code) (*Coupon, *domain.DomainError)

	GetExpiredCoupons() ([]*Coupon, *domain.DomainError)
}
