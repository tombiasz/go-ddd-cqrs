package coupon

import (
	"fmt"
	"go-coupons/src/app/coupons/domain"
	"time"
)

const DefaultCouponExpirationDays = 7

// TODO: create and export Coupon interface to preserve encapsulation
type Coupon struct {
	id          *couponId
	email       *Email
	code        *Code
	description *description
	status      Status
}

func New(
	id string,
	email string,
	code string,
	description string,
	status string,
	expdays uint8,
	activatedAt time.Time,
	expiredAt *time.Time,
	usedAt *time.Time,
) (*Coupon, domain.DomainErrors) {
	cId, errId := NewCouponId(id)
	c, errCode := NewCode(code)
	e, errEmail := NewEmail(email)
	d, errDesc := NewDescription(description)

	errs := domain.CombineDomainErrors(errId, errCode, errEmail, errDesc)

	if errs != nil {
		return nil, errs
	}

	s := NewStatus(
		status,
		expdays,
		activatedAt,
		expiredAt,
		usedAt,
	)

	coupon := &Coupon{
		id:          cId,
		email:       e,
		code:        c,
		description: d,
		status:      s,
	}

	return coupon, nil
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

	status := NewActiveStatus(expirationInDays, timeProvider.Now())

	c := &Coupon{
		id:          GenerateCouponId(),
		email:       _email,
		code:        GenerateCode(),
		description: desc,
		status:      status,
	}

	return c, nil
}

func (c *Coupon) canBeMarkedAsUsed() bool {
	return c.status.Status() == ActiveStatus
}

func (c *Coupon) Use(timeProvider domain.TimeProvider) *domain.DomainError {
	if !c.canBeMarkedAsUsed() {
		return CouponCannotBeMarkedAsUsedErr
	}

	c.status = c.status.Use(timeProvider)

	return nil
}

func (c *Coupon) canBeMarkedAsExpired(timeProvider domain.TimeProvider) bool {
	if !c.status.IsActive() {
		return false
	}

	return c.status.IsExpired(timeProvider)
}

func (c *Coupon) Expire(timeProvider domain.TimeProvider) *domain.DomainError {
	if !c.canBeMarkedAsExpired(timeProvider) {
		return CouponCannotBeMarkedAdExpiredErr
	}

	c.status = c.status.Expire(timeProvider)

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

// TODO: should be moved to coupon
func (c *Coupon) ExpireInDays() uint8 {
	return c.status.ExpiresInDays()
}

func (c *Coupon) ActivatedAt() time.Time {
	return c.status.ActivatedAt()
}

func (c *Coupon) UsedAt() *time.Time {
	return c.status.UsedAt()
}

func (c *Coupon) ExpiredAt() *time.Time {
	return c.status.ExpiredAt()
}

func (c *Coupon) String() string {
	return fmt.Sprintf("Coupon: {\n\tid: %s\n\temail: %s\n\tcode: %s\n\tstatus: %s\n}\n", c.id.value, c.email.address, c.code.value, c.status.Status())
}
