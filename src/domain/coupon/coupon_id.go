package coupon

import (
	"fmt"
	"go-coupons/src/domain"
)

type couponId struct {
	value string
}

func CreateCouponId(id *domain.ID) *couponId {
	return &couponId{id.Value()}
}

func (id *couponId) Value() string {
	return id.value
}

func (id *couponId) String() string {
	return fmt.Sprintf("<CouponID: %s >", id.value)
}
