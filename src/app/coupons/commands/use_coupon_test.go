package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestUseCouponCommand(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("returns error when invalid email or code was passed ", func(t *testing.T) {
		var cmd = &UseCouponCommand{
			Email: "",
			Code:  "",
		}

		var handler = &UseCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   nil,
		}

		err := handler.Execute(cmd)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		var foundEmailErr = lookupError(t, err, coupon.EmailCannotBeEmptyErr)
		var foundCodeErr = lookupError(t, err, coupon.CodeCannotBeEmptyErr)

		if !foundEmailErr {
			t.Errorf("expected %q error but did not received one", coupon.EmailCannotBeEmptyErr)
		}

		if !foundCodeErr {
			t.Errorf("expected %q error but did not received one", coupon.CodeCannotBeEmptyErr)
		}
	})

	t.Run("returns nil when coupon is successfully used", func(t *testing.T) {
		var email = "foo@bar.com"
		var c, _ = coupon.RegisterCoupon(email, "desc", 7, fixedTimeProvider)
		var cmd = &UseCouponCommand{
			Email: c.Email(),
			Code:  c.Code(),
		}

		onGetCouponByEmailAndCode := func(email *coupon.Email, code *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
			if c.Email() != email.Address() {
				t.Errorf("want %q got %q", email, c.Email())
			}

			if c.Code() != code.Value() {
				t.Errorf("want %q got %q", code, c.Code())
			}

			return c, nil
		}

		onSave := func(coupon *coupon.Coupon) *domain.DomainError {
			if coupon != c {
				t.Errorf("want %q got %q", c, coupon)
			}

			return nil
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: onGetCouponByEmailAndCode,
			onGetExpiredCoupons:       nil,
		}

		var handler = &UseCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute(cmd)

		if err != nil {
			t.Errorf("expected to not to receive an error but got one")
		}

		if c.Status() != coupon.UsedStatus {
			t.Errorf("expected to get used coupon but it was not")
		}
	})

	t.Run("returns error if coupon cannot be stored", func(t *testing.T) {
		var email = "foo@bar.com"
		var c, _ = coupon.RegisterCoupon(email, "desc", 7, fixedTimeProvider)
		var cmd = &UseCouponCommand{
			Email: c.Email(),
			Code:  c.Code(),
		}
		var repoErr = &domain.DomainError{"test", "repo error"}

		onGetCouponByEmailAndCode := func(email *coupon.Email, code *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
			return c, nil
		}

		onSave := func(coupon *coupon.Coupon) *domain.DomainError {
			return repoErr
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: onGetCouponByEmailAndCode,
			onGetExpiredCoupons:       nil,
		}

		var handler = &UseCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute(cmd)

		if err == nil {
			t.Errorf("expected to receive an error but did not get one")
		}

		var found = lookupError(t, err, repoErr)

		if !found {
			t.Errorf("expected %q error but did not received one", repoErr)
		}
	})

	t.Run("returns error if coupon cannot be retrieved from repository", func(t *testing.T) {
		var email = "foo@bar.com"
		var c, _ = coupon.RegisterCoupon(email, "desc", 7, fixedTimeProvider)
		var cmd = &UseCouponCommand{
			Email: c.Email(),
			Code:  c.Code(),
		}
		var repoErr = &domain.DomainError{"test", "repo error"}

		onGetCouponByEmailAndCode := func(email *coupon.Email, code *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
			return nil, repoErr
		}

		var fakeRepo = &FakeRepo{
			onSave:                    nil,
			onGetCouponByEmailAndCode: onGetCouponByEmailAndCode,
			onGetExpiredCoupons:       nil,
		}

		var handler = &UseCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute(cmd)

		if err == nil {
			t.Errorf("expected to receive an error but did not get one")
		}

		var found = lookupError(t, err, repoErr)

		if !found {
			t.Errorf("expected %q error but did not received one", repoErr)
		}
	})
}
