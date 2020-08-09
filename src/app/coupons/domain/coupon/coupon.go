package coupon

import (
	"fmt"
	"go-coupons/src/app/coupons/domain"
)

const DefaultCouponExpirationDays = 7

type Coupon struct {
	id          *couponId
	email       *Email
	code        *Code
	description *description
	status      Status
}

func New(
	id *couponId,
	email *Email,
	code *Code,
	description *description,
	status Status,
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
	_email, emailErr := NewEmail(email)
	desc, descErr := NewDescription(description)

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

func (c *Coupon) canBeUsed() bool {
	return c.status.Status() == ActiveStatus
}

func (c *Coupon) Use(timeProvider domain.TimeProvider) *domain.DomainError {
	if !c.canBeUsed() {
		return CouponAlreadyUsedErr
	}

	c.status = c.status.(*activeStatus).Use(timeProvider)

	return nil
}

func (c *Coupon) canBeMarkedAsExpired(timeProvider domain.TimeProvider) bool {
	isActive := c.status.Status() == ActiveStatus

	if !isActive {
		return false
	}

	return c.status.(*activeStatus).isExpired(timeProvider)
}

func (c *Coupon) Expire(timeProvider domain.TimeProvider) *domain.DomainError {
	if !c.canBeMarkedAsExpired(timeProvider) {
		return CouponCannotBeNotExpiredErr
	}

	c.status = c.status.(*activeStatus).Expire()

	return nil
}

func (c *Coupon) Id() string {
	return c.id.Value()
}

func (c *Coupon) Email() string {
	return c.email.address
}

func (c *Coupon) Code() string {
	return c.code.Value()
}

func (c *Coupon) Description() string {
	return c.description.value
}

func (c *Coupon) Status() string {
	return c.status.Status()
}

func (c *Coupon) String() string {
	return fmt.Sprintf("Coupon: {\n\tid: %s\n\temail: %s\n\tcode: %s\n\tstatus: %s\n}\n", c.id.value, c.email.address, c.code.value, c.status.Status())
}
