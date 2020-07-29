package coupon

import "go-coupons/src/app/coupons/domain"

type email struct {
	address string
}

func NewEmail(address string) (*email, *domain.DomainError) {
	if address == "" {
		return nil, EmailCannotBeEmptyErr
	}

	if len(address) < 3 || len(address) > 255 {
		return nil, EmailIsInvalidErr
	}

	return &email{address}, nil
}
