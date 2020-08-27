package notifier

import (
	"fmt"
	"go-coupons/src/app/coupons/domain/coupon"
)

type ExternalApiNotifier struct{}

func (n *ExternalApiNotifier) Notify(coupon *coupon.Coupon) error {
	fmt.Println(fmt.Sprintf("Sending coupon (id: %s) data to external API", coupon.Id()))

	return nil
}
