package coupon

import "go-coupons/src/app/coupons/domain"

type description struct {
	value string
}

func NewDescription(value string) (*description, *domain.DomainError) {
	if value == "" {
		return nil, DescriptionCannotBeEmptyErr
	}

	if len(value) > 200 {
		return nil, DescriptionCannotBeLongerThan200CharsErr
	}

	return &description{value}, nil
}
