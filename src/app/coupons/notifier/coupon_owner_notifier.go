package notifier

import (
	"fmt"
	"go-coupons/src/app/coupons/domain/coupon"
)

type CouponOwnerNotifier struct{}

func (n *CouponOwnerNotifier) Notify(coupon *coupon.Coupon) error {
	fmt.Println(fmt.Sprintf("Sending coupon code %s to email %s", coupon.Code(), coupon.Email()))

	return nil
}
