package coupon

import "fmt"

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

func Create(id, email, code, description, status string) (*coupon, error) {
	e, err1 := createEmail(email)
	d, err2 := createDescription(description)

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	c := &coupon{
		id:          id,
		email:       e,
		code:        code,
		description: d,
		status:      status,
	}

	return c, nil
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
	return c.description.value
}

func (c coupon) Status() string {
	return c.status
}

func (c coupon) String() string {
	return fmt.Sprintf("<Coupon: %s %s (%s)>", c.email, c.code, c.id)
}
