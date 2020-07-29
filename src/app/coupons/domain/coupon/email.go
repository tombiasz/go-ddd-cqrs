package coupon

import "go-coupons/src/app/coupons/domain"

type Email struct {
	address string
}

func NewEmail(address string) (*Email, *domain.DomainError) {
	if address == "" {
		return nil, EmailCannotBeEmptyErr
	}

	if len(address) < 3 || len(address) > 255 {
		return nil, EmailIsInvalidErr
	}

	return &Email{address}, nil
}
