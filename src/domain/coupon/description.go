package coupon

import "go-coupons/src/domain"

type description struct {
	value string
}

func newDescription(value string) *description {
	return &description{value}
}

func CreateDescription(value string) (*description, *domain.DomainError) {
	if value == "" {
		return nil, DescriptionCannotBeEmptyErr
	}

	if len(value) > 200 {
		return nil, DescriptionCannotBeLongerThan200CharsErr
	}

	return &description{value}, nil
}
