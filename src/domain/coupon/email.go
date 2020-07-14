package coupon

import "go-coupons/src/domain"

type email struct {
	address string
}

func newEmail(address string) *email {
	return &email{address}
}

func CreateEmail(address string) (*email, *domain.DomainError) {
	if address == "" {
		return nil, EmailCannotBeEmptyErr
	}

	if len(address) < 3 || len(address) > 255 {
		return nil, EmailIsInvalidErr
	}

	return &email{address}, nil
}
