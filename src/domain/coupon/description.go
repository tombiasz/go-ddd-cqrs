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
		return nil, domain.NewDomainError("description", "description cannot be empty")
	}

	if len(value) > 200 {
		return nil, domain.NewDomainError("description", "description must have less than 200 characters")
	}

	return &description{value}, nil
}
