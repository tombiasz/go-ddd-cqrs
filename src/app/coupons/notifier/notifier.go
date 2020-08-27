package notifier

import (
	"go-coupons/src/app/coupons/domain/coupon"
)

type Notifier interface {
	Notify(*coupon.Coupon) error
}
