package coupon

import "fmt"

type coupon struct {
	id          string
	email       string
	code        string
	description string
	status      string
}

func New(id, email, code, description, status string) *coupon {
	return &coupon{id, email, code, description, status}
}

func (c coupon) Id() string {
	return c.id
}

func (c coupon) Email() string {
	return c.email
}

func (c coupon) Code() string {
	return c.code
}

func (c coupon) Description() string {
	return c.description
}

func (c coupon) Status() string {
	return c.status
}

func (c coupon) String() string {
	return fmt.Sprintf("<Coupon: %s %s (%s)>", c.email, c.code, c.id)
}
