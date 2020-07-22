package coupon

import (
	"fmt"
	"go-coupons/src/app/coupons/domain"
)

const DefaultCouponExpirationDays = 7

type coupon struct {
	id          *couponId
	email       *email
	code        *code
	description *description
	status      status
}

func New(
	id *couponId,
	email *email,
	code *code,
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
	email string,
	description string,
	expirationInDays uint8,
	timeProvider domain.TimeProvider,
) (*coupon, domain.DomainErrors) {
	_email, emailErr := CreateEmail(email)
	desc, descErr := CreateDescription(description)

	err := domain.CombineDomainErrors(emailErr, descErr)

	if err != nil {
		return nil, err
	}

	c := &coupon{
		id:          GenerateCouponId(),
		email:       _email,
		code:        GenerateCode(),
		description: desc,
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
	return c.code.Value()
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
