package coupon

import "go-coupons/src/app/coupons/domain"

type Repository interface {
	Save(coupon *coupon) *domain.DomainError
}
