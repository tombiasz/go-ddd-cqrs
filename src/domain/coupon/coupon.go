package coupon

import "fmt"

type coupon struct {
	id          string
	email       *email
	code        string
	description string
	status      string
}

func New(id, email, code, description, status string) *coupon {
	return &coupon{id, newEmail(email), code, description, status}
}

func Create(id, email, code, description, status string) (*coupon, error) {
	e, err := createEmail(email)

	if err != nil {
		return nil, err
	}

	return &coupon{id, e, code, description, status}, nil
}

func (c coupon) Id() string {
	return c.id
}

func (c coupon) Email() string {
	return c.email.value
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
