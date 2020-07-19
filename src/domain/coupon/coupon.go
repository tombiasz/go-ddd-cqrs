package coupon

import (
	"fmt"
	"go-coupons/src/domain"
)

const DefaultCouponExpirationDays = 7

type coupon struct {
	id          *couponId
	email       *email
	code        string
	description *description
	status      status
}

func New(
	id *couponId,
	email *email,
	code string,
	description *description,
	status status,
) *coupon {
	return &coupon{
		id,
		email,
		code,
		description,
		status,
	}
}

func RegisterCoupon(
	email *email,
	code string,
	description *description,
	expirationInDays uint8,
	timeProvider domain.TimeProvider,
) (*coupon, domain.DomainErrors) {
	var emailErr *domain.DomainError
	if email == nil {
		emailErr = EmailCannotBeNilErr
	}

	var descErr *domain.DomainError
	if description == nil {
		descErr = DescriptionCannotBeNilErr
	}

	err := domain.CombineDomainErrors(emailErr, descErr)
	if err != nil {
		return nil, err
	}

	c := &coupon{
		id:          CreateCouponId(),
		email:       email,
		code:        code,
		description: description,
		status:      CreateActiveStatus(expirationInDays, timeProvider),
	}

	return c, nil
}

func (c coupon) Id() string {
	return c.id.Value()
}

func (c coupon) Email() string {
	return c.email.address
}

func (c coupon) Code() string {
	return c.code
}

func (c coupon) Description() string {
	return c.description.value
}

func (c coupon) Status() string {
	return c.status.Status()
}

func (c coupon) String() string {
	return fmt.Sprintf("<Coupon: %s %s (%s)>", c.email, c.code, c.id)
}
