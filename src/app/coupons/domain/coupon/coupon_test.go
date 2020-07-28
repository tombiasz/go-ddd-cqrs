package coupon

import (
	"go-coupons/src/app/coupons/domain"
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestRegisterCoupon(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("returns error when invalid email was passed ", func(t *testing.T) {
		var desc = "Lorem ipsum dolor sit amet."

		_, err := RegisterCoupon("", desc, 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		found := lookupError(t, err, EmailCannotBeEmptyErr)
		if !found {
			t.Errorf("expected %q error but did not received one", EmailCannotBeEmptyErr)
		}
	})

	t.Run("returns error when invalid value was passed as description", func(t *testing.T) {
		var email = "foo@bar.com"

		_, err := RegisterCoupon(email, "", 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		found := lookupError(t, err, DescriptionCannotBeEmptyErr)
		if !found {
			t.Errorf("expected %q error but did not received one", DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("returns both errors when invalid value was passed as email and description", func(t *testing.T) {
		_, err := RegisterCoupon("", "", 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected errors but did not received any")
		}

		var foundEmailErr = lookupError(t, err, EmailCannotBeEmptyErr)
		var foundDescErr = lookupError(t, err, DescriptionCannotBeEmptyErr)

		if !foundEmailErr {
			t.Errorf("expected %q error but did not received one", EmailCannotBeEmptyErr)
		}

		if !foundDescErr {
			t.Errorf("expected %q error but did not received one", DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("registers a new coupon", func(t *testing.T) {
		var email = "foo@bar.com"
		var desc = "Lorem ipsum dolor sit amet."

		c, _ := RegisterCoupon(email, desc, 7, fixedTimeProvider)

		if c.Id() == "" {
			t.Errorf("expected non empty id but received %q", c.Id())
		}

		if c.Email() != email {
			t.Errorf("expected %q but received %q", email, c.Email())
		}

		if c.Code() == "" {
			t.Errorf("expected non empty code but received %q", c.Code())
		}

		if c.Description() != desc {
			t.Errorf("expected %q but received %q", desc, c.Description())
		}

		if c.Status() != ActiveStatus {
			t.Errorf("expected %q but received %q", ActiveStatus, c.Status())
		}
	})
}

func TestCouponUse(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("marks coupon as used", func(t *testing.T) {
		var email = "foo@bar.com"
		var desc = "Lorem ipsum"

		c, _ := RegisterCoupon(email, desc, 7, fixedTimeProvider)

		c.Use(fixedTimeProvider)

		if c.Status() != UsedStatus {
			t.Errorf("expected %q but received %q", UsedStatus, c.Status())
		}

		s := c.status.(*usedStatus)

		if !s.usedAt.Equal(fakeNow) {
			t.Errorf("expected %q but received %q", UsedStatus, c.Status())
		}
	})

	t.Run("coupon can only be used once", func(t *testing.T) {
		var email = "foo@bar.com"
		var desc = "Lorem ipsum"

		c, _ := RegisterCoupon(email, desc, 7, fixedTimeProvider)

		err := c.Use(fixedTimeProvider)

		if err != nil {
			t.Errorf("did not expect an err but received one %q", err)
		}

		err = c.Use(fixedTimeProvider)

		if err != CouponAlreadyUsedErr {
			t.Errorf("expected %q but received %q", CouponAlreadyUsedErr, err)
		}
	})
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
