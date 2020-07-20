package coupon

import (
	"go-coupons/src/domain"
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestRegisterCoupon(t *testing.T) {
	var fakeNow = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

	t.Run("returns error when nil was passed as email", func(t *testing.T) {
		var desc, _ = CreateDescription("Lorem ipsum dolor sit amet.")

		_, err := RegisterCoupon(nil, desc, 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		found := lookupError(t, err, EmailCannotBeNilErr)
		if !found {
			t.Errorf("expected %q error but did not received one", EmailCannotBeNilErr)
		}
	})

	t.Run("returns error when nil was passed as description", func(t *testing.T) {
		var email, _ = CreateEmail("foo@bar.com")

		_, err := RegisterCoupon(email, nil, 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}

		found := lookupError(t, err, DescriptionCannotBeNilErr)
		if !found {
			t.Errorf("expected %q error but did not received one", DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("returns both errors when nil was passed as email and description", func(t *testing.T) {
		_, err := RegisterCoupon(nil, nil, 7, fixedTimeProvider)

		if err == nil {
			t.Errorf("expected errors but did not received any")
		}

		var foundEmailErr = lookupError(t, err, EmailCannotBeNilErr)
		var foundDescErr = lookupError(t, err, DescriptionCannotBeNilErr)

		if !foundEmailErr {
			t.Errorf("expected %q error but did not received one", EmailCannotBeNilErr)
		}

		if !foundDescErr {
			t.Errorf("expected %q error but did not received one", DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("registers a new coupon", func(t *testing.T) {
		var email, _ = CreateEmail("foo@bar.com")
		var desc, _ = CreateDescription("Lorem ipsum dolor sit amet.")

		c, _ := RegisterCoupon(email, desc, 7, fixedTimeProvider)

		if c.Id() == "" {
			t.Errorf("expected non empty id but received %q", c.Id())
		}

		if c.Email() != email.address {
			t.Errorf("expected %q but received %q", email.address, c.Email())
		}

		if c.Code() == "" {
			t.Errorf("expected non empty code but received %q", c.Code())
		}

		if c.Description() != desc.value {
			t.Errorf("expected %q but received %q", desc.value, c.Description())
		}

		if c.Status() != ActiveStatus {
			t.Errorf("expected %q but received %q", ActiveStatus, c.Status())
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
