package coupon

import (
	"fmt"
	"go-coupons/src/app/coupons/domain"
)

const DefaultCouponExpirationDays = 7

type Coupon struct {
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
) *Coupon {
	return &Coupon{
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
) (*Coupon, domain.DomainErrors) {
	_email, emailErr := CreateEmail(email)
	desc, descErr := CreateDescription(description)

	err := domain.CombineDomainErrors(emailErr, descErr)

	if err != nil {
		return nil, err
	}

	c := &Coupon{
		id:          GenerateCouponId(),
		email:       _email,
		code:        GenerateCode(),
		description: desc,
		status:      CreateActiveStatus(expirationInDays, timeProvider),
	}

	return c, nil
}

func (c Coupon) Id() string {
	return c.id.Value()
}

func (c Coupon) Email() string {
	return c.email.address
}

func (c Coupon) Code() string {
	return c.code.Value()
}

func (c Coupon) Description() string {
	return c.description.value
}

func (c Coupon) Status() string {
	return c.status.Status()
}

func (c Coupon) String() string {
	return fmt.Sprintf("<Coupon: %s %s (%s)>", c.email, c.code, c.id)
}
