package coupon

import "errors"

type email struct {
	value string
}

func newEmail(value string) *email {
	return &email{value}
}

func createEmail(value string) (*email, error) {
	if value == "" {
		return nil, errors.New("email cannot be empty")
	}

	if len(value) < 3 || len(value) > 255 {
		return nil, errors.New("email is invalid")
	}

	return &email{value}, nil
}
