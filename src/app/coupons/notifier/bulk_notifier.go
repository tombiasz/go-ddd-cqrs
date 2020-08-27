package notifier

import "go-coupons/src/app/coupons/domain/coupon"

type CouponBulkNotifier struct {
	notifiers []Notifier
}

func CreateCouponBulkNotifier(notifiers ...Notifier) Notifier {
	return &CouponBulkNotifier{notifiers}
}

func (n *CouponBulkNotifier) Notify(coupon *coupon.Coupon) error {
	for _, n := range n.notifiers {
		n.Notify(coupon)
	}

	return nil
}
