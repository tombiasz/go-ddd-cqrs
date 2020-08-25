package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestExpireCouponsCommandHandler(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("returns error if coupons cannot be retrieved from repository", func(t *testing.T) {
		var repoErr = &domain.DomainError{"test", "repo error"}

		onGetExpiredCoupons := func() ([]*coupon.Coupon, *domain.DomainError) {
			return nil, repoErr
		}

		var fakeRepo = &FakeRepo{
			onSave:                    nil,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       onGetExpiredCoupons,
		}

		var handler = &ExpireCouponsCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute()

		if err == nil {
			t.Errorf("expected to receive an error but did not get one")
		}

		var found = lookupError(t, err, repoErr)

		if !found {
			t.Errorf("expected %q error but did not received one", repoErr)
		}
	})

	t.Run("returns error if coupons cannot be mark as expired", func(t *testing.T) {
		var now = time.Now()
		var fixedTimeProvider = &timeutils.FixedTimeProvider{now}
		var c, _ = coupon.RegisterCoupon(
			"foo@bar.com",
			"desc",
			coupon.DefaultCouponExpirationDays,
			fixedTimeProvider,
		)

		onGetExpiredCoupons := func() ([]*coupon.Coupon, *domain.DomainError) {
			return []*coupon.Coupon{c}, nil
		}

		var fakeRepo = &FakeRepo{
			onSave:                    nil,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       onGetExpiredCoupons,
		}

		var handler = &ExpireCouponsCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute()

		if err == nil {
			t.Errorf("expected to receive an error but did not get one")
		}

		var found = lookupError(t, err, coupon.CouponCannotBeMarkedAdExpiredErr)

		if !found {
			t.Errorf("expected %q error but did not received one", coupon.CouponCannotBeMarkedAdExpiredErr)
		}
	})

	t.Run("returns error if coupons cannot be saved", func(t *testing.T) {
		const expiresInDays = 7
		var now = time.Now()
		var past = now.AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)
		var pastTimeProvider = &timeutils.FixedTimeProvider{past}
		var nowTimeProvider = &timeutils.FixedTimeProvider{now}
		var repoErr = &domain.DomainError{"test", "repo error"}
		var c, _ = coupon.RegisterCoupon(
			"foo@bar.com",
			"desc",
			expiresInDays,
			pastTimeProvider,
		)

		onSave := func(coupon *coupon.Coupon) *domain.DomainError {
			return repoErr
		}

		onGetExpiredCoupons := func() ([]*coupon.Coupon, *domain.DomainError) {
			return []*coupon.Coupon{c}, nil
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       onGetExpiredCoupons,
		}

		var handler = &ExpireCouponsCommandHandler{
			TimeProvider: nowTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute()

		var found = lookupError(t, err, repoErr)
		if !found {
			t.Errorf("expected %q error but did not received one", repoErr)
		}
	})

	t.Run("returns nil after expiring coupons", func(t *testing.T) {
		const expiresInDays = 7
		var now = time.Now()
		var past = now.AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)
		var pastTimeProvider = &timeutils.FixedTimeProvider{past}
		var nowTimeProvider = &timeutils.FixedTimeProvider{now}
		var c, _ = coupon.RegisterCoupon(
			"foo@bar.com",
			"desc",
			expiresInDays,
			pastTimeProvider,
		)

		onSave := func(couponToBeStored *coupon.Coupon) *domain.DomainError {
			if couponToBeStored != c {
				t.Errorf("want %q got %q", couponToBeStored, c)
			}

			if couponToBeStored.Status() != coupon.ExpiredStatus {
				t.Errorf("expected coupon to be expired but it was %q", couponToBeStored.Status())
			}

			return nil
		}

		onGetExpiredCoupons := func() ([]*coupon.Coupon, *domain.DomainError) {
			return []*coupon.Coupon{c}, nil
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       onGetExpiredCoupons,
		}

		var handler = &ExpireCouponsCommandHandler{
			TimeProvider: nowTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute()

		if err != nil {
			t.Errorf("expected not to receive an error but got one %q", err)
		}
	})
}
