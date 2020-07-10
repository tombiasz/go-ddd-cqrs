package coupon

import "errors"

type description struct {
	value string
}

func newDescription(value string) *description {
	return &description{value}
}

func createDescription(value string) (*description, error) {
	if value == "" {
		return nil, errors.New("description cannot be empty")
	}

	if len(value) > 200 {
		return nil, errors.New("description must have less than 200 characters")
	}

	return &description{value}, nil
}
