package commands

import (
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestRegisterCouponCommandHandler(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("returns error when invalid email and description was passed ", func(t *testing.T) {
		var cmd = &RegisterCouponCommand{
			Email:       "",
			Description: "",
		}

		var handler = &RegisterCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   nil,
		}

		err := handler.Execute(cmd)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		var foundEmailErr = lookupError(t, err, coupon.EmailCannotBeEmptyErr)
		var foundDescErr = lookupError(t, err, coupon.DescriptionCannotBeEmptyErr)

		if !foundEmailErr {
			t.Errorf("expected %q error but did not received one", coupon.EmailCannotBeEmptyErr)
		}

		if !foundDescErr {
			t.Errorf("expected %q error but did not received one", coupon.DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("returns nil when coupon is successfully registered", func(t *testing.T) {
		var cmd = &RegisterCouponCommand{
			Email:       "foo@bar.com",
			Description: "Lorem ipsum dolor sit amet.",
		}

		onSave := func(coupon *coupon.Coupon) *domain.DomainError {
			if coupon.Email() != cmd.Email {
				t.Errorf("want %q got %q", cmd.Email, coupon.Email())
			}

			if coupon.Description() != cmd.Description {
				t.Errorf("want %q got %q", cmd.Email, coupon.Email())
			}

			return nil
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       nil,
		}

		var handler = &RegisterCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute(cmd)

		if err != nil {
			t.Errorf("expected to not to receive an error but got one")
		}
	})

	t.Run("returns error if coupon cannot be stored", func(t *testing.T) {
		var cmd = &RegisterCouponCommand{
			Email:       "foo@bar.com",
			Description: "Lorem ipsum dolor sit amet.",
		}

		repoFailure := &domain.DomainError{"test", "repo failure"}

		onSave := func(coupon *coupon.Coupon) *domain.DomainError {
			return repoFailure
		}

		var fakeRepo = &FakeRepo{
			onSave:                    onSave,
			onGetCouponByEmailAndCode: nil,
			onGetExpiredCoupons:       nil,
		}

		var handler = &RegisterCouponCommandHandler{
			TimeProvider: fixedTimeProvider,
			Repository:   fakeRepo,
		}

		err := handler.Execute(cmd)

		if err == nil {
			t.Errorf("expected to receive an error but did not get one")
		}

		var found = lookupError(t, err, repoFailure)

		if !found {
			t.Errorf("expected %q error but did not received one", repoFailure)
		}
	})
}
