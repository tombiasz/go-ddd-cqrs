package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	"testing"
)

type FakeRepo struct {
	onSave func(*coupon.Coupon) *domain.DomainError

	onGetCouponByEmailAndCode func(*coupon.Email, *coupon.Code) (*coupon.Coupon, *domain.DomainError)

	onGetExpiredCoupons func() ([]*coupon.Coupon, *domain.DomainError)
}

func (r *FakeRepo) Save(coupon *coupon.Coupon) *domain.DomainError {
	return r.onSave(coupon)
}

func (r *FakeRepo) GetCouponByEmailAndCode(email *coupon.Email, code *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
	return r.onGetCouponByEmailAndCode(email, code)
}
func (r *FakeRepo) GetExpiredCoupons() ([]*coupon.Coupon, *domain.DomainError) {
	return r.onGetExpiredCoupons()
}

func lookupError(t *testing.T, errs domain.DomainErrors, err error) bool {
	t.Helper()

	var found = false

	for _, e := range errs {
		if e == err {
			found = true
			break
		}
	}

	return found
}

type FakeNotifier struct {
	onNotify func(*coupon.Coupon) error
}

func (n *FakeNotifier) Notify(coupon *coupon.Coupon) error {
	return n.onNotify(coupon)
}
