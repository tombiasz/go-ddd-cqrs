package coupon

import (
	"fmt"
)

type coupon struct {
	id          string
	email       *email
	code        string
	description *description
	status      string
}

func New(id, email, code, description, status string) *coupon {
	return &coupon{
		id:          id,
		email:       newEmail(email),
		code:        code,
		description: newDescription(description),
		status:      status,
	}
}

func Create(id string, email *email, code string, description *description, status string) *coupon {
	c := &coupon{
		id:          id,
		email:       email,
		code:        code,
		description: description,
		status:      status,
	}

	return c
}

func (c coupon) Id() string {
	return c.id
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
	return c.status
}

func (c coupon) String() string {
	return fmt.Sprintf("<Coupon: %s %s (%s)>", c.email, c.code, c.id)
}
